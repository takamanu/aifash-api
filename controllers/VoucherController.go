package controllers

import (
	"context"
	"curdusers/configs"
	"curdusers/models"
	"io"
	"net/http"
	"strconv"
	"time"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

// CreateVoucher creates a new voucher.
func CreateVoucher(c echo.Context, client *storage.Client, bucketName *string) error {

	// Parse and validate the request
	request := models.Voucher{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	// Create the voucher in the database
	image, err := c.FormFile("voucher_url_image")
	if err != nil {
		log.Errorf("Failed to get the voucher image file: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Voucher image upload failed")
	}

	uniqueFilename := uuid.NewString()

	ctx := context.Background()
	wc := client.Bucket(*bucketName).Object(uniqueFilename).NewWriter(ctx)
	defer wc.Close()

	src, err := image.Open()
	if err != nil {
		log.Errorf("Failed to open the image file: %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to process image")
	}
	defer src.Close()

	if _, err = io.Copy(wc, src); err != nil {
		log.Errorf("Failed to copy the image to the bucket: %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to upload image")
	}

	voucherPriceString := c.FormValue("voucher_price")
	voucherPrice, err := strconv.ParseUint(voucherPriceString, 10, 0)
	if err != nil {
		log.Errorf("Failed to convert VoucherPrice to a uint: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Voucher Price")
	}

	currentTime := time.Now()
	duration := 7 * 24 * time.Hour
	expirationDate := currentTime.Add(duration)

	request.VoucherName = c.FormValue("voucher_name")
	request.VoucherCode = c.FormValue("voucher_code")
	request.VoucherValue = uint(voucherPrice)
	request.ExpirationDate = expirationDate

	// Set the ImageUrl to the URL of the uploaded image in the existing bucket
	request.VoucherImageUrl = "https://storage.googleapis.com/" + *bucketName + "/" + uniqueFilename

	// Create the fashion item in your database
	if err := configs.DB.Create(&request).Error; err != nil {
		log.Errorf("Failed to create voucher: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Voucher created successfully",
		"data":    request,
	})
}

// GetVouchers retrieves all vouchers.
func GetVouchers(c echo.Context) error {
	var request []models.Voucher

	// if err := configs.DB.Preload("Voucher").Find(&request).Error; err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	// }

	if err := configs.DB.Find(&request).Error; err != nil {
		log.Errorf("Failed to get vouchers: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if len(request) == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "There is no vouchers.",
			// "data":    request,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Get all vouchers",
		"data":    request,
	})
}

// GetVoucherByID retrieves a voucher by its ID.
func GetVoucherByID(c echo.Context) error {
	itemID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Voucher ID")
	}

	var request models.Voucher
	// Query the database to find the fashion item by ID
	if err := configs.DB.First(&request, itemID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Voucher not found")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success fetch voucher",
		"data":    request,
	})
}

// UpdateVoucher updates a voucher by its ID.
func UpdateVoucher(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	voucherID := c.Param("id")

	var voucher models.Voucher
	if err := db.First(&voucher, voucherID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"error": "Voucher not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	// Parse and validate the request
	updateData := new(models.Voucher)
	if err := c.Bind(updateData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	// Update the voucher in the database
	if err := db.Model(&voucher).Updates(updateData).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, voucher)
}

// DeleteVoucher deletes a voucher by its ID.
func DeleteVoucher(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	voucherID := c.Param("id")

	var voucher models.Voucher
	if err := db.First(&voucher, voucherID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"error": "Voucher not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	// Delete the voucher from the database
	if err := db.Delete(&voucher).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.NoContent(http.StatusNoContent)
}

// UpdateVoucherController updates a voucher by its ID.
func UpdateVoucherController(c echo.Context) error {
	voucherID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Voucher ID")
	}

	// Fetch the existing voucher from the database
	var existingVoucher models.Voucher
	if err := configs.DB.First(&existingVoucher, voucherID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "Voucher not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Parse the updated data from the request
	updatedVoucher := models.Voucher{}
	if err := c.Bind(&updatedVoucher); err != nil {
		log.Errorf("Failed to bind request: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Update the fields of the existing voucher with the new data
	existingVoucher.VoucherName = updatedVoucher.VoucherName
	existingVoucher.VoucherCode = updatedVoucher.VoucherCode
	existingVoucher.VoucherValue = updatedVoucher.VoucherValue
	// You can update other fields here as well

	// Save the updated voucher to the database
	if err := configs.DB.Save(&existingVoucher).Error; err != nil {
		log.Errorf("Failed to update voucher: %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update voucher")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Voucher updated successfully",
		"data":    existingVoucher,
	})
}

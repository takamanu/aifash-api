package controllers

import (
	"curdusers/configs"
	"curdusers/models" // Import your models package
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func CreateUserVoucher(c echo.Context) error {
	// Parse and validate the request
	request := models.UserVoucher{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	// Set the UserVoucher as not used
	request.Status = false

	// Fetch the user from the database based on the provided UserID
	var user models.User

	if err := configs.DB.First(&user, request.UserID).Error; err != nil {
		log.Errorf("Failed to fetch user from the database: %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user voucher")
	}

	// Fetch the voucher from the database based on the provided VoucherID
	var voucher models.Voucher
	if err := configs.DB.First(&voucher, request.VoucherID).Error; err != nil {
		log.Errorf("Failed to fetch voucher from the database: %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user voucher")
	}

	// Calculate the points to deduct from the user
	pointsToDeduct := uint(voucher.VoucherValue)

	// Check if the user has enough points to buy the voucher
	if user.Points < pointsToDeduct {
		return echo.NewHTTPError(http.StatusBadRequest, "Insufficient points to buy the voucher")
	}

	// Update the user's points
	user.Points -= pointsToDeduct
	request.Voucher = voucher

	// Set the User in the UserVoucher

	tx := configs.DB.Begin()
	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	tx.Commit()

	// Create the UserVoucher in the database
	if err := configs.DB.Create(&request).Error; err != nil {
		log.Errorf("Failed to create voucher for user: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Voucher created successfully",
		"data":    request,
	})
}

// MarkVoucherAsUsed marks a UserVoucher as used when a user clicks "use" on a voucher.
func MarkVoucherAsUsed(c echo.Context) error {
	userVoucherID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Voucher Users")
	}

	var userVoucher models.UserVoucher
	if err := configs.DB.First(&userVoucher, userVoucherID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Voucher user not found")
	}

	// Toggle the UserVoucher's Used field
	userVoucher.Status = !userVoucher.Status

	// Update the UserVoucher in the database
	if err := configs.DB.Save(&userVoucher).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Fetch the associated User and Voucher data
	var user models.User
	if err := configs.DB.First(&user, userVoucher.UserID).Error; err != nil {
		log.Errorf("Failed to fetch user from the database: %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch user data")
	}

	var voucher models.Voucher
	if err := configs.DB.First(&voucher, userVoucher.VoucherID).Error; err != nil {
		log.Errorf("Failed to fetch voucher from the database: %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch voucher data")
	}

	userVoucher.Voucher = voucher

	// Return the UserVoucher along with associated User and Voucher data
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Voucher Toggle",
		"data":    userVoucher,
	})
}

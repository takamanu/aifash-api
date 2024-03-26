package controllers

import (
	"context"
	"curdusers/configs"
	"curdusers/models"
	"io"
	"net/http"
	"strconv"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log" // Import your GORM models package
	"gorm.io/gorm"
)

// CreateFashionItemController handles the creation of a new fashion item.
func CreateFashionItemController(c echo.Context, client *storage.Client, bucketName *string) error {
	// Parse request and validate input
	request := models.Fashion{}
	if err := c.Bind(&request); err != nil {
		log.Errorf("Failed to bind request: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Handle file upload
	image, err := c.FormFile("fashion_url_image")
	if err != nil {
		log.Errorf("Failed to get the image file: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Image upload failed")
	}

	// Generate a unique filename using a UUID
	uniqueFilename := uuid.NewString()

	// Upload the image to the existing Google Cloud Storage bucket
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

	// Convert the UserID to an integer
	userIDStr := c.FormValue("user_id")
	userID, err := strconv.ParseUint(userIDStr, 10, 0)
	if err != nil {
		log.Errorf("Failed to convert UserID to a uint: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid UserID")
	}

	fashionPointsStr := c.FormValue("fashion_points")
	fashionPoints, err := strconv.ParseUint(fashionPointsStr, 10, 0)
	if err != nil {
		log.Errorf("Failed to convert fashion points to a uint: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid fashion points")
	}

	// Set the UserID in the request as a uint
	request.UserID = uint(userID)

	request.ItemType = c.FormValue("fashion_name")
	request.ItemPoints = int(fashionPoints)
	request.ItemStatus = "on_process"

	// request.UploadTimestamp = time.Now().String()

	// Set the ImageUrl to the URL of the uploaded image in the existing bucket
	request.ImageUrl = "https://storage.googleapis.com/" + *bucketName + "/" + uniqueFilename

	// Create the fashion item in your database
	if err := configs.DB.Create(&request).Error; err != nil {
		log.Errorf("Failed to create fashion: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Fashion item created successfully",
		"data":    request,
	})
}

// GetFashionItemsController retrieves all fashion items.
func GetFashionItemsController(c echo.Context) error {
	var request []models.Fashion

	if err := configs.DB.Find(&request).Error; err != nil {
		log.Errorf("Failed to get fashion: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if len(request) == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "There is no fashion.",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Get all fashion",
		"data":    request,
	})
}

// GetFashionItemByIDController retrieves a fashion item by its ID.
func GetFashionItemByIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Errorf("Invalid id: %s", c.Param("id"))
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid id")
	}

	var user models.Fashion
	if err := configs.DB.First(&user, id).Error; err != nil {
		log.Errorf("Failed to get fashion with id %d: %s", id, err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success fetch data!",
		"data":    user,
	})
}

func ChooseResponseFashion(c echo.Context) error {
	fashionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Fashion ID")
	}

	var fashion models.Fashion
	if err := configs.DB.First(&fashion, fashionID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Fashion not found")
	}

	// Check if the fashion status is already "accepted" or "denied"
	if fashion.ItemStatus == "accepted" || fashion.ItemStatus == "denied" {
		return echo.NewHTTPError(http.StatusBadRequest, "Fashion status is already finalized and cannot be updated.")
	}

	var user models.User
	if err := configs.DB.First(&user, fashion.UserID).Error; err != nil {
		log.Errorf("Failed to fetch user from the database: %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch user data")
	}

	// Get the desired status from the request
	desiredStatus := c.FormValue("status")

	// Check if the desired status is valid
	switch desiredStatus {
	case "on_process", "accepted", "denied":
		fashion.ItemStatus = models.FashionStatus(desiredStatus)
	default:
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid status")
	}

	// Calculate the points to be added or deducted
	pointsToChange := uint(fashion.ItemPoints)

	if fashion.ItemStatus == "accepted" {
		user.Points += pointsToChange
	}

	// Add error handling for database updates
	tx := configs.DB.Begin()
	if err := tx.Save(&fashion).Error; err != nil {
		tx.Rollback()
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	tx.Commit()

	// Fetch the associated User and Voucher data

	// Return the fashion along with associated User and Voucher data
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Fashion Change Request",
		"data":    fashion,
	})
}

// UpdateFashionItemController updates a fashion item by its ID.
func UpdateFashionItemController(c echo.Context) error {
	itemID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Fashion ID")
	}

	// Fetch the existing fashion item from the database
	var existingFashion models.Fashion
	if err := configs.DB.First(&existingFashion, itemID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "Fashion not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Parse the updated data from the request
	updatedFashion := models.Fashion{}
	if err := c.Bind(&updatedFashion); err != nil {
		log.Errorf("Failed to bind request: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Check if ItemStatus is true in the existing item, and if so, prevent updating ItemPoints

	// Update the fields of the existing fashion item with the new data
	existingFashion.ItemType = updatedFashion.ItemType

	// Only update ItemPoints if ItemStatus is false

	if existingFashion.ItemStatus != "accepted" || existingFashion.ItemStatus != "denied" {
		existingFashion.ItemPoints = updatedFashion.ItemPoints
	}

	// Save the updated fashion item to the database
	if err := configs.DB.Save(&existingFashion).Error; err != nil {
		log.Errorf("Failed to update fashion: %s", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update fashion")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Fashion item updated successfully",
		"data":    existingFashion,
	})
}

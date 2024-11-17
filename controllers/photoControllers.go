package controllers

import (
	"fmt"
	"mygram/database"
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreatePhoto(c *gin.Context) {
	var db = database.GetDB()
	var photo models.Photo

	// Parse the JSON request and populate the Photo struct
	err := c.ShouldBindJSON(&photo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}

	// Validate request on create photo context.
	var CreatePhotoRules models.CreatePhotoRules
	CreatePhotoRules.Title = photo.Title
	CreatePhotoRules.Caption = photo.Caption
	CreatePhotoRules.PhotoURL = photo.PhotoURL

	validate := validator.New()
	err = validate.Struct(CreatePhotoRules)
	fmt.Println(err) // Logging error on console... just because.

	if err != nil {
		// Extracting validation errors
		errorDetails := make(map[string]string)
		for _, validationErr := range err.(validator.ValidationErrors) {
			errorDetails[validationErr.Field()] = fmt.Sprintf("Validation failed on '%s' tag", validationErr.Tag())
		}

		// Return error.
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": errorDetails,
		})
		return
	}

	// Get userdata from JWT.
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	// Create photo.
	photoInput := models.Photo{
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoURL: photo.PhotoURL,
		UserID:   userID,
	}

	err = db.Create(&photoInput).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"error":  err,
		})
		return
	}

	// Return success response.
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   photoInput,
	})
}

func UpdatePhoto(c *gin.Context) {
	var db = database.GetDB()
	var photo models.Photo

	// Parse the JSON request and populate the Photo struct
	err := c.ShouldBindJSON(&photo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}

	// Validate request on update photo context.
	var UpdatePhotoRules models.UpdatePhotoRules
	UpdatePhotoRules.ID = photo.ID
	UpdatePhotoRules.Title = photo.Title
	UpdatePhotoRules.Caption = photo.Caption
	UpdatePhotoRules.PhotoURL = photo.PhotoURL

	validate := validator.New()
	err = validate.Struct(UpdatePhotoRules)
	fmt.Println(err) // Logging error on console... just because.

	if err != nil {
		// Extracting validation errors
		errorDetails := make(map[string]string)
		for _, validationErr := range err.(validator.ValidationErrors) {
			errorDetails[validationErr.Field()] = fmt.Sprintf("Validation failed on '%s' tag", validationErr.Tag())
		}

		// Return error.
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": errorDetails,
		})
		return
	}

	// Check if the photo ID exists in database.
	var existingPhoto models.Photo
	err = db.First(&existingPhoto, "id = ?", photo.ID).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  "Photo is not found in database.",
		})
		return
	}

	// Get userdata from JWT.
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	// Check if the photo belongs to user.
	if existingPhoto.UserID != userID {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "This photo does not belong to current user.",
		})
		return
	}

	// Update photo.
	existingPhoto.PhotoURL = photo.PhotoURL
	existingPhoto.Caption = photo.Caption
	existingPhoto.Title = photo.Title

	if err := db.Save(&existingPhoto).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}

	// Return success response.
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   existingPhoto,
	})
}

func DeletePhoto(c *gin.Context) {
	var db = database.GetDB()
	var photo models.Photo

	// Parse the JSON request and populate the Photo struct
	err := c.ShouldBindJSON(&photo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}

	// Validate request on update photo context.
	var DeletePhotoRules models.DeletePhotoRules
	DeletePhotoRules.ID = photo.ID

	validate := validator.New()
	err = validate.Struct(DeletePhotoRules)
	fmt.Println(err) // Logging error on console... just because.

	if err != nil {
		// Extracting validation errors
		errorDetails := make(map[string]string)
		for _, validationErr := range err.(validator.ValidationErrors) {
			errorDetails[validationErr.Field()] = fmt.Sprintf("Validation failed on '%s' tag", validationErr.Tag())
		}

		// Return error.
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": errorDetails,
		})
		return
	}

	// Check if the photo ID exists in database.
	var existingPhoto models.Photo
	err = db.First(&existingPhoto, "id = ?", photo.ID).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  "Photo is not found in database.",
		})
		return
	}

	// Get userdata from JWT.
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	// Check if the photo belongs to user.
	if existingPhoto.UserID != userID {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "This photo does not belong to current user.",
		})
		return
	}

	// Delete photo.
	err = db.Delete(&existingPhoto).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}

	// Return success response.
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Photo has been deleted.",
	})
}

func GetAllPhotos(c *gin.Context) {
	// Declare variables.
	var db = database.GetDB()
	var photos []models.Photo

	// Get userdata from JWT.
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	// Retrieve all photos belonging to user.
	err := db.Where("user_id = ?", userID).Find(&photos).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}

	// Return photos.
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   photos,
	})
}

func GetOnePhoto(c *gin.Context) {
	// Declare variables.
	var db = database.GetDB()
	var photo models.Photo
	var photoID uint = 0

	// Get userdata from JWT.
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	// Get photoID from the URL.
	photoIDStr := c.Query("photo_id")
	photoID64, err := strconv.ParseUint(photoIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Please make sure to add '?photo_id=1' to the URL, this is a GET API.",
		})
		return
	}
	photoID = uint(photoID64)

	// Retrieve the specific photo belonging to user.
	err = db.Where("id = ? AND user_id = ?", uint(photoID), userID).First(&photo).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"error":  "Photo is either not found, or does not belong to user.",
		})
		return
	}

	// Return photos.
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   photo,
	})
}

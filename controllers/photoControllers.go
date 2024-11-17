package controllers

import (
	"fmt"
	"mygram/database"
	"mygram/models"
	"net/http"

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

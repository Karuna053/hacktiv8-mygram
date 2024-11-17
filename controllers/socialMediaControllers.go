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

func CreateSocialMedia(c *gin.Context) {
	var db = database.GetDB()
	var socialMedia models.SocialMedia

	// Parse the JSON request and populate the SocialMedia struct
	err := c.ShouldBindJSON(&socialMedia)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}

	// Validate request on create social media context.
	var CreateSocialMediaRules models.CreateSocialMediaRules
	CreateSocialMediaRules.Name = socialMedia.Name
	CreateSocialMediaRules.SocialMediaURL = socialMedia.SocialMediaURL

	validate := validator.New()
	err = validate.Struct(CreateSocialMediaRules)
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

	// Create social media.
	socialMediaInput := models.SocialMedia{
		Name:           socialMedia.Name,
		SocialMediaURL: socialMedia.SocialMediaURL,
		UserID:         userID,
	}

	err = db.Create(&socialMediaInput).Error
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
		"data":   socialMediaInput,
	})
}

func UpdateSocialMedia(c *gin.Context) {
	var db = database.GetDB()
	var socialMedia models.SocialMedia

	// Parse the JSON request and populate the Social Media struct
	err := c.ShouldBindJSON(&socialMedia)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}

	// Validate request on update social media context.
	var UpdateSocialMediaRules models.UpdateSocialMediaRules
	UpdateSocialMediaRules.ID = socialMedia.ID
	UpdateSocialMediaRules.Name = socialMedia.Name
	UpdateSocialMediaRules.SocialMediaURL = socialMedia.SocialMediaURL

	validate := validator.New()
	err = validate.Struct(UpdateSocialMediaRules)
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

	// Check if the social media ID exists in database.
	var existingSocialMedia models.SocialMedia
	err = db.First(&existingSocialMedia, "id = ?", socialMedia.ID).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  "Social Media is not found in database.",
		})
		return
	}

	// Get userdata from JWT.
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	// Check if the social media belongs to user.
	if existingSocialMedia.UserID != userID {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "This social media does not belong to current user.",
		})
		return
	}

	// Update social media.
	existingSocialMedia.Name = socialMedia.Name
	existingSocialMedia.SocialMediaURL = socialMedia.SocialMediaURL

	if err := db.Save(&existingSocialMedia).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}

	// Return success response.
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   existingSocialMedia,
	})
}

func DeleteSocialMedia(c *gin.Context) {
	var db = database.GetDB()
	var socialMedia models.SocialMedia

	// Parse the JSON request and populate the Social Media struct
	err := c.ShouldBindJSON(&socialMedia)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}

	// Validate request on update social media context.
	var DeleteSocialMediaRules models.DeleteSocialMediaRules
	DeleteSocialMediaRules.ID = socialMedia.ID

	validate := validator.New()
	err = validate.Struct(DeleteSocialMediaRules)
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

	// Check if the social media ID exists in database.
	var existingSocialMedia models.SocialMedia
	err = db.First(&existingSocialMedia, "id = ?", socialMedia.ID).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  "Social Media is not found in database.",
		})
		return
	}

	// Get userdata from JWT.
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	// Check if the social media belongs to user.
	if existingSocialMedia.UserID != userID {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "This social media does not belong to current user.",
		})
		return
	}

	// Delete socialMedia.
	err = db.Delete(&existingSocialMedia).Error
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
		"message": "Social Media has been deleted.",
	})
}

func GetAllSocialMedias(c *gin.Context) {
	// Declare variables.
	var db = database.GetDB()
	var socialMedias []models.SocialMedia

	// Get userdata from JWT.
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	// Retrieve all social medias belonging to user.
	err := db.Where("user_id = ?", userID).Find(&socialMedias).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}

	// Return social medias.
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   socialMedias,
	})
}

func GetOneSocialMedia(c *gin.Context) {
	// Declare variables.
	var db = database.GetDB()
	var socialMedia models.SocialMedia
	var socialMediaID uint = 0

	// Get userdata from JWT.
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	// Get socialMediaID from the URL.
	socialMediaIDStr := c.Query("social_media_id")
	socialMediaID64, err := strconv.ParseUint(socialMediaIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Please make sure to add '?social_media_id=1' to the URL, this is a GET API.",
		})
		return
	}
	socialMediaID = uint(socialMediaID64)

	// Retrieve the specific social media belonging to user.
	err = db.Where("id = ? AND user_id = ?", uint(socialMediaID), userID).First(&socialMedia).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"error":  "Social Media is either not found, or does not belong to user.",
		})
		return
	}

	// Return social medias.
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   socialMedia,
	})
}

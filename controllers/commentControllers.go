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

func CreateComment(c *gin.Context) {
	var db = database.GetDB()
	var comment models.Comment

	// Parse the JSON request and populate the Comment struct
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}

	// Validate request on create comment context.
	var CreateCommentRules models.CreateCommentRules
	CreateCommentRules.PhotoID = comment.PhotoID
	CreateCommentRules.Message = comment.Message

	validate := validator.New()
	err = validate.Struct(CreateCommentRules)
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
	err = db.First(&existingPhoto, "id = ?", comment.PhotoID).Error
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

	// Create comment.
	commentInput := models.Comment{
		UserID:  userID,
		PhotoID: existingPhoto.ID,
		Message: comment.Message,
	}

	err = db.Create(&commentInput).Error
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
		"data":   commentInput,
	})
}

func UpdateComment(c *gin.Context) {
	var db = database.GetDB()
	var comment models.Comment

	// Parse the JSON request and populate the Photo struct
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}

	// Validate request on update photo context.
	var UpdateCommentRules models.UpdateCommentRules
	UpdateCommentRules.ID = comment.ID
	UpdateCommentRules.Message = comment.Message

	validate := validator.New()
	err = validate.Struct(UpdateCommentRules)
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

	// Check if comment ID exists in database.
	var existingComment models.Comment
	err = db.First(&existingComment, "id = ?", comment.ID).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  "Comment not found in database.",
		})
		return
	}

	// Get userdata from JWT.
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	// Check if the comment belongs to user.
	if existingComment.UserID != userID {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "This comment does not belong to current user.",
		})
		return
	}

	// Update comment.
	existingComment.Message = comment.Message

	if err := db.Save(&existingComment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}

	// Return success response.
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   existingComment,
	})
}

func DeleteComment(c *gin.Context) {
	var db = database.GetDB()
	var comment models.Comment

	// Parse the JSON request and populate the Comment struct
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}

	// Validate request on update comment context.
	var DeleteCommentRules models.DeleteCommentRules
	DeleteCommentRules.ID = comment.ID

	validate := validator.New()
	err = validate.Struct(DeleteCommentRules)
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

	// Check if comment ID exists in database.
	var existingComment models.Comment
	err = db.First(&existingComment, "id = ?", comment.ID).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  "Comment not found in database.",
		})
		return
	}

	// Get userdata from JWT.
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	// Check if the comment belongs to user.
	if existingComment.UserID != userID {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "This comment does not belong to current user.",
		})
		return
	}

	// Delete comment.
	err = db.Delete(&existingComment).Error
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
		"message": "Comment has been deleted.",
	})
}

func GetAllComments(c *gin.Context) {
	// Declare variables.
	var db = database.GetDB()
	var comments []models.Comment

	// Get userdata from JWT.
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	// Retrieve all comments belonging to user.
	err := db.Where("user_id = ?", userID).Find(&comments).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}

	// Return comments.
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   comments,
	})
}

func GetOneComment(c *gin.Context) {
	// Declare variables.
	var db = database.GetDB()
	var comment models.Comment
	var commentID uint = 0

	// Get userdata from JWT.
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	// Get commentID from the URL.
	commentIDStr := c.Query("comment_id")
	commentID64, err := strconv.ParseUint(commentIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Please make sure to add '?comment_id=1' to the URL, this is a GET API.",
		})
		return
	}
	commentID = uint(commentID64)

	// Retrieve the specific comment belonging to user.
	err = db.Where("id = ? AND user_id = ?", uint(commentID), userID).First(&comment).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"error":  "Comment is either not found, or does not belong to user.",
		})
		return
	}

	// Return comments.
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   comment,
	})
}

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

// func UserRegister(c *gin.Context) {
// 	// Declare variables.
// 	db := database.GetDB()
// 	contentType := helpers.GetContentType(c)
// 	var User models.User

// 	// Parse the JSON request and populate the User struct
// 	if contentType == appJSON {
// 		c.ShouldBindJSON(&User)
// 	} else {
// 		c.ShouldBind(&User)
// 	}

// 	// Validate user on login context.
// 	var UserRegisterRules models.UserRegisterRules
// 	UserRegisterRules.Username = User.Username
// 	UserRegisterRules.Email = User.Email
// 	UserRegisterRules.Password = User.Password
// 	UserRegisterRules.Age = User.Age

// 	validate := validator.New()
// 	err := validate.Struct(UserRegisterRules)
// 	fmt.Println(err) // Logging error on console... just because.

// 	if err != nil {
// 		// Extracting validation errors
// 		errorDetails := make(map[string]string)
// 		for _, validationErr := range err.(validator.ValidationErrors) {
// 			errorDetails[validationErr.Field()] = fmt.Sprintf("Validation failed on '%s' tag", validationErr.Tag())
// 		}

// 		// Return error.
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  "fail",
// 			"message": errorDetails,
// 		})
// 		return
// 	}

// 	// Encrypt password.
// 	User.Password = helpers.HashPass(User.Password)

// 	// Create User.
// 	err = db.Debug().Create(&User).Error
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  "fail",
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	// Return a successful message.
// 	c.JSON(http.StatusCreated, gin.H{
// 		"status": "success",
// 		"data": []gin.H{
// 			{
// 				"id":       User.ID,
// 				"username": User.Username,
// 				"email":    User.Email,
// 				"password": User.Password, // Make sure this is encrypted.
// 				"age":      User.Age,
// 			},
// 		},
// 	})
// }

// func UserLogin(c *gin.Context) {
// 	// Declare variables.
// 	db := database.GetDB()
// 	contentType := helpers.GetContentType(c)
// 	var User models.User

// 	// Parse the JSON request and populate the User struct
// 	if contentType == appJSON {
// 		c.ShouldBindJSON(&User)
// 	} else {
// 		c.ShouldBind(&User)
// 	}

// 	// Validate on login context.
// 	var UserLoginRules models.UserLoginRules
// 	UserLoginRules.Username = User.Username
// 	UserLoginRules.Email = User.Email
// 	UserLoginRules.Password = User.Password

// 	validate := validator.New()
// 	err := validate.Struct(UserLoginRules)
// 	fmt.Println(err) // Logging error on console... just because.

// 	if err != nil {
// 		// Extracting validation errors
// 		errorDetails := make(map[string]string)
// 		for _, validationErr := range err.(validator.ValidationErrors) {
// 			errorDetails[validationErr.Field()] = fmt.Sprintf("Validation failed on '%s' tag", validationErr.Tag())
// 		}

// 		// Return error.
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  "fail",
// 			"message": errorDetails,
// 		})
// 		return
// 	}

// 	// Perform login attempt.
// 	var password string = User.Password

// 	err = db.Debug().Where("email = ?", User.Email).Take(&User).Error           // Check if email matches.
// 	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password)) // Check if password matches.
// 	if err != nil || !comparePass {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"status":  "fail",
// 			"message": "Invalid email/password",
// 		})
// 		return
// 	}

// 	// Generate login token.
// 	token := helpers.GenerateToken(User.ID, User.Email)

// 	// Return success message.
// 	c.JSON(http.StatusOK, gin.H{
// 		"status": "success",
// 		"token":  token,
// 	})
// }

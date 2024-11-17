package controllers

import "github.com/gin-gonic/gin"

// var (
// 	appJSON = "application/json"
// )

func CreatePhoto(c *gin.Context) {

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

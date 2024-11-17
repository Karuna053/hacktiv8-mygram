package routers

import (
	"mygram/controllers"
	"mygram/middlewares"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userRouter := router.Group("/api/user")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	photoRouter := router.Group("/api/photo")
	{
		photoRouter.Use(middlewares.Authentication())
		// photoRouter.GET("/create", controllers.GetOnePhoto)
		// photoRouter.GET("/create", controllers.GetAllPhotos)
		photoRouter.POST("/create", controllers.CreatePhoto)
		photoRouter.PUT("/update", controllers.UpdatePhoto)
		photoRouter.DELETE("/delete", controllers.DeletePhoto)
	}

	return router
}

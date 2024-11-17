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
		photoRouter.GET("/get-one", controllers.GetOnePhoto)
		photoRouter.GET("/get-all", controllers.GetAllPhotos)
		photoRouter.POST("/create", controllers.CreatePhoto)
		photoRouter.PUT("/update", controllers.UpdatePhoto)
		photoRouter.DELETE("/delete", controllers.DeletePhoto)
	}

	socialMediaRouter := router.Group("/api/social-media")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.GET("/get-one", controllers.GetOneSocialMedia)
		socialMediaRouter.GET("/get-all", controllers.GetAllSocialMedias)
		socialMediaRouter.POST("/create", controllers.CreateSocialMedia)
		socialMediaRouter.PUT("/update", controllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/delete", controllers.DeleteSocialMedia)
	}

	// commentRouter := router.Group("/api/comment")
	// {
	// 	commentRouter.Use(middlewares.Authentication())
	// 	commentRouter.GET("/get-one", controllers.GetOneComment)
	// 	commentRouter.GET("/get-all", controllers.GetAllComments)
	// 	commentRouter.POST("/create", controllers.CreateComment)
	// 	commentRouter.PUT("/update", controllers.UpdateComment)
	// 	commentRouter.DELETE("/delete", controllers.DeleteComment)
	// }

	return router
}

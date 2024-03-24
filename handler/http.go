package handler

import (
	"final-project/handler/comment"
	"final-project/handler/photo"
	"final-project/handler/social_media"
	"final-project/handler/user"
	"final-project/helpers"
	"final-project/middlewares"

	"github.com/gin-gonic/gin"
)

func NewHTTPServer(userHandler *user.Handler, socialMediaHandler *social_media.Handler, photoHandler *photo.Handler, commentHandler *comment.Handler) {
	r := gin.Default()
	port := helpers.GetEnv("PORT")

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", userHandler.RegisterUser)
		userRouter.POST("/login", userHandler.LoginUser)
		userRouter.Use(middlewares.Authentication())
		userRouter.PUT("/", userHandler.UpdateUser)
		userRouter.DELETE("/", userHandler.DeleteUser)
	}

	socialMediaRouter := r.Group("/socialmedias")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.GET("/", socialMediaHandler.GetAllSocialMedia)
		socialMediaRouter.GET("/:id", socialMediaHandler.GetSocialMediaById)
		socialMediaRouter.POST("/", socialMediaHandler.CreateSocialMedia)
		socialMediaRouter.PUT("/:id", middlewares.Authorization("social_media"), socialMediaHandler.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:id", middlewares.Authorization("social_media"), socialMediaHandler.DeleteSocialMedia)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.GET("/", photoHandler.GetAllPhoto)
		photoRouter.GET("/:id", photoHandler.GetPhotoById)
		photoRouter.POST("/", photoHandler.CreatePhoto)
		photoRouter.PUT("/:id", middlewares.Authorization("photo"), photoHandler.UpdatePhoto)
		photoRouter.DELETE("/:id", middlewares.Authorization("photo"), photoHandler.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.GET("/", commentHandler.GetAllComment)
		commentRouter.GET("/:id", commentHandler.GetCommentById)
		commentRouter.POST("/", commentHandler.CreateComment)
		commentRouter.PUT("/:id", middlewares.Authorization("comment"), commentHandler.UpdateComment)
		commentRouter.DELETE("/:id", middlewares.Authorization("comment"), commentHandler.DeleteComment)
	}

	r.Run(":" + port)
}

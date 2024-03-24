package main

import (
	"final-project/database"
	"final-project/handler"
	commentHandler "final-project/handler/comment"
	photoHandler "final-project/handler/photo"
	socialMediaHandler "final-project/handler/social_media"
	userHandler "final-project/handler/user"
	"final-project/repository/comment"
	photo "final-project/repository/photo"
	socialMedia "final-project/repository/social_media"
	"final-project/repository/user"
	commentSvc "final-project/service/comment"
	photoSvc "final-project/service/photo"
	socialMediaSvc "final-project/service/social_media"
	userSvc "final-project/service/user"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
	db, err := database.NewDatabase()
	if err != nil {
		panic(err)
	}

	//user
	userRepo := user.NewRepository(db)
	userService := userSvc.NewService(userRepo)
	userHandler := userHandler.NewHandler(userService)

	//social media
	socialMediaRepo := socialMedia.NewRepository(db)
	socialMediaService := socialMediaSvc.NewService(socialMediaRepo)
	socialMediaHandler := socialMediaHandler.NewHandler(socialMediaService)

	//photo
	photoRepo := photo.NewRepository(db)
	photoService := photoSvc.NewService(photoRepo)
	photoHandler := photoHandler.NewHandler(photoService)

	//photo
	commentRepo := comment.NewRepository(db)
	commentService := commentSvc.NewService(commentRepo)
	commentHandler := commentHandler.NewHandler(commentService)

	handler.NewHTTPServer(userHandler, socialMediaHandler, photoHandler, commentHandler)

}

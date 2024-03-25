package handler

import (
	_ "final-project/docs"
	"final-project/infra/config"
	"final-project/infra/database"

	"final-project/repository/comment_repository/comment_pg"
	"final-project/repository/photo_repository/photo_pg"
	"final-project/repository/social_media_repository/social_media_pg"
	"final-project/repository/user_repository/user_pg"

	"final-project/service/auth_service"
	"final-project/service/comment_service"
	"final-project/service/photo_service"
	"final-project/service/social_media_service"
	"final-project/service/user_service"

	"github.com/gin-gonic/gin"

	swaggoFile "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartApp() {

	config.LoadEnv()

	database.InitializeDatabase()
	db := database.GetInstanceDatabaseConnection()

	userRepo := user_pg.NewUserRepository(db)
	userService := user_service.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	photoRepo := photo_pg.NewPhotoRepository(db)
	photoService := photo_service.NewPhotoService(photoRepo)
	photoHandler := NewPhotoHandler(photoService)

	commentRepo := comment_pg.NewCommentRepository(db)
	commentService := comment_service.NewCommentService(commentRepo, photoRepo)
	commentHandler := NewCommentHandler(commentService)

	socialMediaRepo := social_media_pg.NewSocialMediaRepository(db)
	socialMediaService := social_media_service.NewSocialMediaService(socialMediaRepo)
	socialMediaHandler := NewSocialMediasHandler(socialMediaService)

	authService := auth_service.NewAuthService(userRepo, photoRepo, commentRepo, socialMediaRepo)

	app := gin.Default()

	// swagger
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggoFile.Handler))

	// routing
	users := app.Group("users")

	{
		users.POST("/register", userHandler.Register)
		users.POST("/login", userHandler.Login)
		users.PUT("", authService.Authentication(), userHandler.Update)
		users.DELETE("", authService.Authentication(), userHandler.Delete)
	}

	photos := app.Group("photos")

	{

		photos.Use(authService.Authentication())

		photos.POST("", photoHandler.AddPhoto)
		photos.GET("", photoHandler.GetPhotos)
		photos.PUT("/:photoId", authService.AuthorizationPhoto(), photoHandler.UpdatePhoto)
		photos.DELETE("/:photoId", authService.AuthorizationPhoto(), photoHandler.DeletePhoto)
	}

	comments := app.Group("comments")

	{
		comments.POST("", authService.Authentication(), commentHandler.AddComment)
		comments.GET("", authService.Authentication(), commentHandler.GetComments)
		comments.PUT("/:commentId", authService.Authentication(), authService.AuthorizationComment(), commentHandler.UpdateComment)
		comments.DELETE("/:commentId", authService.Authentication(), authService.AuthorizationComment(), commentHandler.DeleteComment)
	}

	socialMedias := app.Group("socialmedias")

	{
		socialMedias.POST("", authService.Authentication(), socialMediaHandler.AddSocialMedia)
		socialMedias.GET("", authService.Authentication(), socialMediaHandler.GetSocialMedias)
		socialMedias.PUT("/:socialMediaId", authService.Authentication(), authService.AuthorizationSocialMedia(), socialMediaHandler.UpdateSocialMedia)
		socialMedias.DELETE("/:socialMediaId", authService.Authentication(), authService.AuthorizationSocialMedia(), socialMediaHandler.DeleteSocialMedia)
	}

	app.Run(":" + config.AppConfig().Port)
}

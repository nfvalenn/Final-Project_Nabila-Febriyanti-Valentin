package main

import (
	"MY-GRAM/controllers"
	"MY-GRAM/lib"
	"MY-GRAM/middleware"
	"MY-GRAM/models"
	"MY-GRAM/repository"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi koneksi database
	db, err := lib.InitDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Migrasi model ke database
	db.AutoMigrate(&models.User{}, &models.Photo{}, &models.SocialMedia{}, &models.Comment{})

	// Inisialisasi repository
	userRepo := repository.NewUserRepository(db)
	photoRepo := repository.NewPhotoRepository(db)
	commentRepo := repository.NewCommentRepository(db)
	socialMediaRepo := repository.NewSocialMediaRepository(db)

	// Inisialisasi controller
	userController := controllers.NewUserController(*userRepo)
	photoController := controllers.NewPhotoController(*photoRepo)
	commentController := controllers.NewCommentController(*commentRepo)
	socialMediaController := controllers.NewSocialMediaController(*socialMediaRepo)

	// Inisialisasi router Gin
	router := gin.Default()

	// Middleware untuk autentikasi
	router.Use(middleware.AuthMiddleware)

	// Routes untuk UserController
	router.POST("/users", userController.RegisterUser)
	router.POST("/login", userController.LoginUser) // Endpoint login yang berbeda
	router.PUT("/users/:userId", userController.UpdateUser)
	router.DELETE("/users/:userId", userController.DeleteUser)

	// Routes untuk PhotoController
	router.POST("/photos", photoController.CreatePhoto)
	router.GET("/photos", photoController.GetPhotos)
	router.PUT("/photos/:photoId", photoController.UpdatePhoto)
	router.DELETE("/photos/:photoId", photoController.DeletePhoto)

	// Routes untuk CommentController
	router.POST("/comments", commentController.CreateComment)
	router.GET("/comments", commentController.GetComments)
	router.PUT("/comments/:commentId", commentController.UpdateComment)
	router.DELETE("/comments/:commentId", commentController.DeleteComment)

	// Routes untuk SocialMediaController
	router.POST("/socialmedias", socialMediaController.CreateSocialMedia)
	router.GET("/socialmedias", socialMediaController.GetSocialMedias)
	router.PUT("/socialmedias/:socialMediaId", socialMediaController.UpdateSocialMedia)
	router.DELETE("/socialmedias/:socialMediaId", socialMediaController.DeleteSocialMedia)

	// Jalankan server
	router.Run(":8080")
}

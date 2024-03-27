package controllers

import (
	"MY-GRAM/models"
	"MY-GRAM/repository"
	"MY-GRAM/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PhotoController struct {
	PhotoRepository repository.PhotoRepository
}

func NewPhotoController(photoRepository repository.PhotoRepository) *PhotoController {
	return &PhotoController{
		PhotoRepository: photoRepository,
	}
}

func (pc *PhotoController) CreatePhoto(c *gin.Context) {
	var photo models.Photo
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Mendapatkan userID dari token JWT yang terautentikasi
	tokenString := c.GetHeader("Authorization")
	_, err := utils.VerifyJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	photo.UserID = photo.User.ID

	createdPhoto, err := pc.PhotoRepository.CreatePhoto(&photo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create photo"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": createdPhoto})
}

func (pc *PhotoController) GetPhotos(c *gin.Context) {
	photos, err := pc.PhotoRepository.GetPhotos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch photos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": photos})
}

func (pc *PhotoController) UpdatePhoto(c *gin.Context) {
	photoID, err := strconv.ParseUint(c.Param("photoId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid photo ID"})
		return
	}

	var updatedPhoto models.Photo
	if err := c.ShouldBindJSON(&updatedPhoto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Mendapatkan userID dari token JWT yang terautentikasi
	tokenString := c.GetHeader("Authorization")
	userID, err := utils.VerifyJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Periksa apakah pengguna yang memperbarui foto adalah pemiliknya
	photo, err := pc.PhotoRepository.GetPhotoByID(uint(photoID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get photo"})
		return
	}
	if photo.UserID != userID.UserID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	updatedPhoto.ID = uint(photoID)
	err = pc.PhotoRepository.UpdatePhoto(uint(photoID), &updatedPhoto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo updated successfully"})
}

func (pc *PhotoController) DeletePhoto(c *gin.Context) {
	photoID, err := strconv.ParseUint(c.Param("photoId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid photo ID"})
		return
	}

	// Mendapatkan userID dari token JWT yang terautentikasi
	tokenString := c.GetHeader("Authorization")
	userID, err := utils.VerifyJWT(tokenString)
	if err != nil {
		// Token tidak valid atau tidak ada
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Periksa apakah pengguna yang menghapus foto adalah pemiliknya
	photo, err := pc.PhotoRepository.GetPhotoByID(uint(photoID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get photo"})
		return
	}
	if photo.UserID != userID.UserID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	err = pc.PhotoRepository.DeletePhoto(uint(photoID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo deleted successfully"})
}

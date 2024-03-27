package irepository

import "MY-GRAM/models"

type PhotoRepository interface {
	CreatePhoto(photo *models.Photo) (*models.Photo, error)
	GetPhotos() ([]*models.Photo, error)
	GetPhotoByID(photoID uint) (*models.Photo, error)
	UpdatePhoto(photoID uint, updatedPhoto *models.Photo) error
	DeletePhoto(photoID uint) error
}

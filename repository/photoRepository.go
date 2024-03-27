package repository

import (
	"MY-GRAM/models"

	"github.com/jinzhu/gorm"
)

type PhotoRepository struct {
	DB *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *PhotoRepository {
	return &PhotoRepository{DB: db}
}

func (pr *PhotoRepository) CreatePhoto(photo *models.Photo) (*models.Photo, error) {
	if err := pr.DB.Create(photo).Error; err != nil {
		return nil, err
	}
	return photo, nil
}

func (pr *PhotoRepository) GetPhotos() ([]models.Photo, error) {
	var photos []models.Photo
	if err := pr.DB.Find(&photos).Error; err != nil {
		return nil, err
	}
	return photos, nil
}

func (p *PhotoRepository) GetPhotoByID(photoID uint) (*models.Photo, error) {
	var photo models.Photo
	if err := p.DB.First(&photo, photoID).Error; err != nil {
		return nil, err
	}
	return &photo, nil
}

func (p *PhotoRepository) UpdatePhoto(photoID uint, updatedPhoto *models.Photo) error {
	var exitingPhoto models.Photo
	if err := p.DB.First(&exitingPhoto, photoID).Error; err != nil {
		return err
	}

	exitingPhoto.Title = updatedPhoto.Title
	exitingPhoto.Caption = updatedPhoto.Caption
	exitingPhoto.PhotoURL = updatedPhoto.PhotoURL

	return p.DB.Save(&exitingPhoto).Error
}

func (p *PhotoRepository) DeletePhoto(photoID uint) error {
	return p.DB.Delete(&models.Photo{}, photoID).Error
}

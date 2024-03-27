package repository

import (
	"MY-GRAM/models"

	"github.com/jinzhu/gorm"
)

type SocialMediaRepository struct {
	DB *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) *SocialMediaRepository {
	return &SocialMediaRepository{
		DB: db,
	}
}

func (smr *SocialMediaRepository) CreateSocialMedia(socialMedia *models.SocialMedia) (*models.SocialMedia, error) {
	if err := smr.DB.Create(socialMedia).Error; err != nil {
		return nil, err
	}
	return socialMedia, nil
}

func (smr *SocialMediaRepository) GetSocialMedias() ([]models.SocialMedia, error) {
	var socialMedias []models.SocialMedia
	if err := smr.DB.Find(&socialMedias).Error; err != nil {
		return nil, err
	}
	return socialMedias, nil
}

func (smr *SocialMediaRepository) GetSocialMediaByID(id uint) (*models.SocialMedia, error) {
	var socialMedia models.SocialMedia
	if err := smr.DB.First(&socialMedia, id).Error; err != nil {
		return nil, err
	}
	return &socialMedia, nil
}

func (smr *SocialMediaRepository) UpdateSocialMedia(updatedSocialMedia *models.SocialMedia) error {
	return smr.DB.Save(updatedSocialMedia).Error
}

func (smr *SocialMediaRepository) DeleteSocialMedia(id uint) error {
	return smr.DB.Delete(&models.SocialMedia{}, id).Error
}

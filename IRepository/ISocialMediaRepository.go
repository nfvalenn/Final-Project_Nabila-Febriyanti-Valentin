package irepository

import "MY-GRAM/models"

type ISocialMediaRepository interface {
	CreateSocialMedia(socialMedia *models.SocialMedia) error
	GetSocialMediaByID(socialMediaID uint) (*models.SocialMedia, error)
	UpdateSocialMedia(socialMediaID uint, updateSocialMedia *models.SocialMedia) error
	DeleteSocialMedia(socialMediaID uint) error
}

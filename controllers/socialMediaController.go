package controllers

import (
	"MY-GRAM/models"
	"MY-GRAM/repository"
	"MY-GRAM/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SocialMediaController struct {
	SocialMediaRepository repository.SocialMediaRepository
}

func NewSocialMediaController(socialMediaRepository repository.SocialMediaRepository) *SocialMediaController {
	return &SocialMediaController{
		SocialMediaRepository: socialMediaRepository,
	}
}

func (smc *SocialMediaController) CreateSocialMedia(c *gin.Context) {
	var socialMedia models.SocialMedia
	if err := c.ShouldBindJSON(&socialMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := getUserIDFromToken(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	socialMedia.UserID = userID

	createdSocialMedia, err := smc.SocialMediaRepository.CreateSocialMedia(&socialMedia)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create social media"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": createdSocialMedia})
}

func (smc *SocialMediaController) GetSocialMedias(c *gin.Context) {
	socialMedias, err := smc.SocialMediaRepository.GetSocialMedias()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch social medias"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": socialMedias})
}

func (smc *SocialMediaController) UpdateSocialMedia(c *gin.Context) {
	socialMediaID, err := strconv.ParseUint(c.Param("socialMediaId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid social media ID"})
		return
	}

	var updatedSocialMedia models.SocialMedia
	if err := c.ShouldBindJSON(&updatedSocialMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := getUserIDFromToken(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	socialMedia, err := smc.SocialMediaRepository.GetSocialMediaByID(uint(socialMediaID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get social media"})
		return
	}
	if socialMedia.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	updatedSocialMedia.ID = uint(socialMediaID)
	err = smc.SocialMediaRepository.UpdateSocialMedia(&updatedSocialMedia)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update social media"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Social media updated successfully"})
}

func (smc *SocialMediaController) DeleteSocialMedia(c *gin.Context) {
	socialMediaID, err := strconv.ParseUint(c.Param("socialMediaId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid social media ID"})
		return
	}

	userID, err := getUserIDFromToken(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	socialMedia, err := smc.SocialMediaRepository.GetSocialMediaByID(uint(socialMediaID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get social media"})
		return
	}
	if socialMedia.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	err = smc.SocialMediaRepository.DeleteSocialMedia(uint(socialMediaID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete social media"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Social media deleted successfully"})
}

func getUserIDFromToken(req *http.Request) (uint, error) {
	tokenString := req.Header.Get("Authorization")
	if tokenString == "" {
		return 0, errors.New("Token not found")
	}
	userID, err := utils.VerifyJWT(tokenString)
	if err != nil {
		return 0, err
	}
	return userID.UserID, nil
}

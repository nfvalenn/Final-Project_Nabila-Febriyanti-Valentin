package repository

import (
	"MY-GRAM/models"

	"github.com/jinzhu/gorm"
)

type CommentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{DB: db}
}

func (cr *CommentRepository) CreateComment(comment *models.Comment) (*models.Comment, error) {
	if err := cr.DB.Create(comment).Error; err != nil {
		return nil, err
	}

	return comment, nil
}

func (c *CommentRepository) GetCommentByID(commentID uint) (*models.Comment, error) {
	var comment models.Comment
	if err := c.DB.First(&comment, commentID).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func (c *CommentRepository) UpdateComment(commentID uint, updateComment *models.Comment) error {
	var exitingComment models.Comment
	if err := c.DB.First(&exitingComment, commentID).Error; err != nil {
		return err
	}

	exitingComment.Message = updateComment.Message

	return c.DB.Save(&exitingComment).Error
}

func (c *CommentRepository) DeleteComment(commentID uint) error {
	return c.DB.Delete(&models.Comment{}, commentID).Error
}

package irepository

import "MY-GRAM/models"

type ICommentRepository interface {
	CreateComment(comment *models.Comment) error
	GetCommentByID(commentID uint) (*models.Comment, error)
	UpdateComment(commentID uint, updatedComment *models.Comment) error
	DeleteComment(commentID uint) error
}

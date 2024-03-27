package controllers

import (
	"MY-GRAM/models"
	"MY-GRAM/repository"
	"MY-GRAM/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	CommentRepository repository.CommentRepository
}

func NewCommentController(commentRepository repository.CommentRepository) *CommentController {
	return &CommentController{
		CommentRepository: commentRepository,
	}
}

func (cc *CommentController) CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
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
	comment.UserID = userID.UserID

	createdComment, err := cc.CommentRepository.CreateComment(&comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": createdComment})
}

func (cc *CommentController) GetComments(c *gin.Context) {
	// Mendapatkan token dari header Authorization
	tokenString := c.GetHeader("Authorization")

	// Mendapatkan userID dari token JWT yang terautentikasi
	userID, err := utils.VerifyJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	comments, err := cc.CommentRepository.GetCommentByID(userID.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comments"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": comments})
}

func (cc *CommentController) UpdateComment(c *gin.Context) {
	commentID, err := strconv.ParseUint(c.Param("commentId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	var updatedComment models.Comment
	if err := c.ShouldBindJSON(&updatedComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Memperbarui komentar
	err = cc.CommentRepository.UpdateComment(uint(commentID), &updatedComment)
	if err != nil {
		if err.Error() == "unauthorized" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update comment"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment updated successfully"})
}

func (cc *CommentController) DeleteComment(c *gin.Context) {
	commentID, err := strconv.ParseUint(c.Param("commentId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	// Mendapatkan token dari header Authorization
	tokenString := c.GetHeader("Authorization")

	// Mendapatkan userID dari token JWT yang terautentikasi
	userID, err := utils.VerifyJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Periksa apakah pengguna yang menghapus komentar adalah pemiliknya
	comment, err := cc.CommentRepository.GetCommentByID(uint(commentID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get comment"})
		return
	}
	if comment.UserID != userID.UserID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	err = cc.CommentRepository.DeleteComment(uint(commentID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}

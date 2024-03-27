package irepository

import "MY-GRAM/models"

type IUserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(userID uint) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(userID uint) error
}

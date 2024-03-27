package repository

import (
	"MY-GRAM/models"
	"errors"

	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	if err := ur.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := u.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	err := u.DB.First(&user, userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) UpdateUser(user *models.User) (*models.User, error) {
	if err := ur.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) DeleteUser(user *models.User) error {
	if err := ur.DB.Delete(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) AuthenticateUser(email, password string) (*models.User, error) {
	var user models.User
	result := r.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}

	// Verifikasi password
	if err := models.VerifyPassword(user.Password, password); err != nil {
		return nil, err
	}

	return &user, nil
}

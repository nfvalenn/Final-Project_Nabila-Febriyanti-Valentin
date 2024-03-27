package models

import (
	"errors"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"unique" json:"username"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	Age       int       `gorm:"not null" json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) Validate() error {
	// Validasi email
	if !IsValidEmail(u.Email) {
		return errors.New("invalid email format")
	}

	// Validasi username tidak boleh kosong
	if IsEmpty(u.Username) {
		return errors.New("username cannot be empty")
	}

	// Validasi password minimal 6 karakter
	if len(u.Password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}

	// Validasi usia minimal 8 tahun
	if u.Age < 8 {
		return errors.New("age must be at least 8")
	}

	return nil
}

// IsValidEmail checks if an email is valid
func IsValidEmail(email string) bool {
	// Regular expression for basic email validation
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// IsEmpty checks if a string is empty or not
func IsEmpty(str string) bool {
	return len(str) == 0
}

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword memverifikasi apakah password sesuai dengan hash yang diberikan
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

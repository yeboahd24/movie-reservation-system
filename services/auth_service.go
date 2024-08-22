package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/yeboahd24/movie-reservation-system/models"
	"github.com/yeboahd24/movie-reservation-system/utils"
	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		DB: db,
	}
}

func (as *AuthService) SignUp(user models.User) (*models.User, error) {
	user.ID = uuid.New()
	user.Role = models.RegularUserRole
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword

	if err := as.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (as *AuthService) Login(credentials models.Credentials) (*models.User, string, error) {
	var user models.User
	if err := as.DB.Where("email = ?", credentials.Email).First(&user).Error; err != nil {
		return nil, "", err
	}

	if !utils.CheckPasswordHash(credentials.Password, user.Password) {
		return nil, "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWTToken(user.ID.String(), string(user.Role))
	if err != nil {
		return nil, "", err
	}

	return &user, token, nil
}

func (as *AuthService) PromoteToAdmin(userID string) error {
	var user models.User
	if err := as.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return err
	}

	user.Role = models.AdminRole
	if err := as.DB.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

package repositories

import (
	"example/latihan-golang-gin/config"
	"example/latihan-golang-gin/models"
)

type UserRepository struct{}

func (r *UserRepository) CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := config.DB.Find(&users).Error
	return users, err
}

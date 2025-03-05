package services

import (
	"example/latihan-golang-gin/models"
	"example/latihan-golang-gin/repositories"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: repo}
}

func (s *UserService) RegisterUser(name, email string) (*models.User, error) {
	user := &models.User{Name: name, Email: email}
	err := s.userRepo.CreateUser(user)
	return user, err
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return s.userRepo.GetUserByEmail(email)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.GetAllUsers()
}

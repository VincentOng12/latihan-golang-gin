package usecase

import (
	"example/latihan-golang-gin/internal/entity"
	"example/latihan-golang-gin/internal/repository"
)

type UserUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) *UserUseCase {
	return &UserUseCase{userRepo: repo}
}

func (uc *UserUseCase) GetUsers() ([]entity.User, error) {
	return uc.userRepo.GetAllUsers()
}

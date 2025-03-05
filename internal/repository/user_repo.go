package repository

import "example/latihan-golang-gin/internal/entity"

type UserRepository interface {
	GetAllUsers() ([]entity.User, error)
}

type userRepo struct{}

func NewUserRepo() UserRepository {
	return &userRepo{}
}

func (r *userRepo) GetAllUsers() ([]entity.User, error) {
	users := []entity.User{
		{ID: 1, Name: "John Doe", Email: "john@example.com"},
		{ID: 2, Name: "Jane Doe", Email: "jane@example.com"},
	}
	return users, nil
}

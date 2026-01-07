package repository

import "github.com/hans312456/go-gin-project/internal/model"

type UserRepository interface {
	FindByID(id string) (*model.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) FindByID(id string) (*model.User, error) {
	// mock example
	return &model.User{
		ID:   id,
		Name: "Hans",
	}, nil
}

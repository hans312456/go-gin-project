package service

import (
	"errors"

	"github.com/hans312456/go-gin-project/internal/repository"

	"github.com/hans312456/go-gin-project/internal/model"
)

type UserService interface {
	GetByID(id string) (*model.User, error)
	Create(name string) (*model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService() UserService {
	return &userService{
		repo: repository.NewUserRepository(),
	}
}

func (s *userService) GetByID(id string) (*model.User, error) {
	if id == "" {
		return nil, errors.New("id cannot be empty")
	}

	return s.repo.FindByID(id)
}

func (s *userService) Create(name string) (*model.User, error) {
	return &model.User{
		ID:   "1",
		Name: name,
	}, nil
}

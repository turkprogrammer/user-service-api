package service

import (
	"awesomeProject/internal/core/domain"
	"awesomeProject/internal/core/ports"
	"github.com/google/uuid"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(name string) (domain.User, error) {
	user := domain.User{
		ID:   uuid.New().String(),
		Name: name,
	}
	err := s.repo.Save(user)
	return user, err
}

func (s *UserService) GetUser(id string) (domain.User, error) {
	return s.repo.Get(id)
}

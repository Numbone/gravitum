package service

import (
	"context"

	"github.com/Numbone/user-api/internal/model"
	"github.com/Numbone/user-api/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, user *model.User) error {
	return s.repo.CreateUser(ctx, user)
}

func (s *UserService) GetUser(ctx context.Context, id int) (*model.User, error) {
	return s.repo.GetUser(ctx, id)
}

func (s *UserService) UpdateUser(ctx context.Context, user *model.User) error {
	return s.repo.UpdateUser(ctx, user)
}

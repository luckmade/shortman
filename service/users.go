package service

import (
	"context"
	"github.com/luckmade/shorter-url/models"
)

type UsersService struct {
	repo *models.UsersRepository
}

func NewUserService(repo *models.UsersRepository) models.UsersService {
	return &UsersService{
		repo: repo,
	}
}

// CreateUser implements models.UsersService.
func (u *UsersService) CreateUser(ctx context.Context, user *models.User) error {
	panic("unimplemented")
}

// DeleteUser implements models.UsersService.
func (u *UsersService) DeleteUser(ctx context.Context, id string) error {
	panic("unimplemented")
}

// GetUserById implements models.UsersService.
func (u *UsersService) GetUserById(ctx context.Context, id string) (*models.User, error) {
	panic("unimplemented")
}

// UpdateUser implements models.UsersService.
func (u *UsersService) UpdateUser(ctx context.Context, user *models.User) error {
	panic("unimplemented")
}


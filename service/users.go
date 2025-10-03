package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/luckmade/shorter-url/models"
	"golang.org/x/crypto/bcrypt"
)

type UsersService struct {
	repo models.UsersRepository
}

func NewUserService(repo models.UsersRepository) models.UsersService {
	return &UsersService{
		repo: repo,
	}
}

// CreateUser implements models.UsersService.
func (u *UsersService) CreateUser(ctx context.Context, data *models.UserData) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	now := time.Now().UTC()
	user := &models.User{
		Id:           uuid.NewString(),
		Name:         data.Name,
		Email:        data.Email,
		Password:     models.Password{Hash: passwordHash},
		CreatedAt:    now,
		LastModified: now,
	}

	return u.repo.Create(ctx, user)
}

// DeleteUser implements models.UsersService.
func (u *UsersService) DeleteUser(ctx context.Context, id string) error {
	return u.repo.Delete(ctx, id)
}

// GetUserById implements models.UsersService.
func (u *UsersService) GetUserById(ctx context.Context, id string) (*models.User, error) {
	return u.repo.Get(ctx, id)
}

// UpdateUser implements models.UsersService.
func (u *UsersService) UpdateUser(ctx context.Context, user *models.User) error {
	return u.repo.Update(ctx, user)
}

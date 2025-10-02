package postgres

import (
	"context"
	"database/sql"

	"github.com/luckmade/shortman/models"
)

type UsersRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) models.UsersRepository {
	return &UsersRepository{
		db: db,
	}
}

// Create implements models.UserRepository.
func (u *UsersRepository) Create(ctx context.Context, user *models.User) error {
	panic("unimplemented")
}

// Delete implements models.UserRepository.
func (u *UsersRepository) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// Get implements models.UserRepository.
func (u *UsersRepository) Get(ctx context.Context, id string) (*models.User, error) {
	panic("unimplemented")
}

// Update implements models.UserRepository.
func (u *UsersRepository) Update(ctx context.Context, user *models.User) error {
	panic("unimplemented")
}


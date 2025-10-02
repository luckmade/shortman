package postgres

import (
	"context"
	"database/sql"

	"github.com/luckmade/shortman/models"
)

type LinksRepository struct {
	db *sql.DB
}

func NewLinksRepository(db *sql.DB) models.LinksRepository {
	return &LinksRepository{
		db: db,
	}
}

// Create implements models.LinksRepository.
func (l *LinksRepository) Create(ctx context.Context, link *models.Link) error {
	panic("unimplemented")
}

// Delete implements models.LinksRepository.
func (l *LinksRepository) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// Get implements models.LinksRepository.
func (l *LinksRepository) Get(ctx context.Context, id string) (*models.Link, error) {
	panic("unimplemented")
}

// Update implements models.LinksRepository.
func (l *LinksRepository) Update(ctx context.Context, link *models.Link) error {
	panic("unimplemented")
}

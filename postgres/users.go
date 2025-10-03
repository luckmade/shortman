package postgres

import (
	"context"
	"database/sql"
	"log"

	"github.com/luckmade/shorter-url/models"
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
	query := `INSERT INTO users(id, name, email, password_hash, created_at, last_modified)
	VALUES($1, NULLIF($2,''), NULLIF($3,''), $4, $5, $6);`

	_, err := u.db.ExecContext(ctx, query, user.Id, user.Name, user.Email, user.Password.Hash, user.CreatedAt, user.LastModified)
	if err != nil {
		log.Printf("failed to insert user: %v", err)
		return err
	}

	return nil
}

// Delete implements models.UserRepository.
func (u *UsersRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1;`
	result, err := u.db.ExecContext(ctx, query, id)
	if err != nil {
		log.Printf("failed to delete user: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("failed to get rows affected: %v", err)
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// Get implements models.UserRepository.
func (u *UsersRepository) Get(ctx context.Context, id string) (*models.User, error) {
	query := `SELECT id, name, email, password_hash, created_at, last_modified FROM users
	WHERE id = $1;`
	row := u.db.QueryRowContext(ctx, query, id)

	user := &models.User{}
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password.Hash, &user.CreatedAt, &user.LastModified)
	if err != nil {
		log.Printf("failed to fetch user: %v", err)
		return nil, err

	}

	return user, nil
}

// Update implements models.UserRepository.
func (u *UsersRepository) Update(ctx context.Context, user *models.User) error {
	query := `UPDATE users
        SET name = NULLIF($2, ''),
            email = NULLIF($3, ''),
            password_hash = $4,
            last_modified = $5
        WHERE id = $1;`

	result, err := u.db.ExecContext(ctx, query, user.Id, user.Name, user.Email, user.Password.Hash, user.LastModified)
	if err != nil {
		log.Printf("failed to update user: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("failed to get rows affected: %v", err)
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

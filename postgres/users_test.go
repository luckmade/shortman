package postgres_test

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand/v2"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/luckmade/shortman/models"
	"github.com/luckmade/shortman/postgres"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func openTestDB(t *testing.T) *sql.DB {
	conn, err := sql.Open("pgx", os.Getenv("TEST_DATABASE_URL"))
	if err != nil {
		t.Errorf("failed to open database: %v", err)
	}

	err = conn.PingContext(context.Background())
	if err != nil {
		t.Errorf("failed to connect to db: %v", err)
	}

	return conn
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.IntN(len(charset))]
	}

	return string(b)
}

func generateValidUser(t *testing.T) *models.User {
	hash, err := bcrypt.GenerateFromPassword([]byte("newpass"), bcrypt.DefaultCost)
	if err != nil {
		t.Errorf("failed to generate password hash: %v", err)
	}
	name := fmt.Sprintf("user%s", generateString(7))
	user := &models.User{
		Id:           uuid.NewString(),
		Name:         name,
		Email:        fmt.Sprintf("%s@mail.com", name),
		Password:     models.Password{Hash: hash},
		CreatedAt:    time.Now().UTC(),
		LastModified: time.Now().UTC(),
	}

	return user
}

func TestUsersRepository_Create(t *testing.T) {
	repo := postgres.NewUserRepository(openTestDB(t))

	user := generateValidUser(t)
	testcases := []struct {
		name      string
		expectErr bool
		user      *models.User
	}{
		{
			name:      "valid user",
			expectErr: false,
			user:      user,
		},
		{
			name:      "invalid user: missing required field",
			expectErr: true,
			user: &models.User{
				Id:   uuid.NewString(),
				Name: "Becky Lynch",
			},
		},
		{
			name:      "invalid user: existing id",
			expectErr: true,
			user: &models.User{
				Id:    user.Id,
				Name:  "Jonny Wave",
				Email: fmt.Sprintf("%s@mail.com", generateString(7)),
			},
		},
	}

	for _, tc := range testcases {
		err := repo.Create(context.Background(), tc.user)
		if tc.expectErr {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}

func TestUsersRepository_Get(t *testing.T) {
	repo := postgres.NewUserRepository(openTestDB(t))

	user := generateValidUser(t)
	err := repo.Create(context.Background(), user)
	require.NoError(t, err)

	t.Run("get existing user", func(t *testing.T) {
		got, err := repo.Get(context.Background(), user.Id)
		require.NoError(t, err)
		require.NotNil(t, got)
		require.Equal(t, user.Id, got.Id)
		require.Equal(t, user.Name, got.Name)
		require.Equal(t, user.Email, got.Email)
	})

	t.Run("get non-existing user", func(t *testing.T) {
		got, err := repo.Get(context.Background(), uuid.NewString())
		require.Error(t, err)
		require.Nil(t, got)
	})
}

func TestUsersRepository_Update(t *testing.T) {
	repo := postgres.NewUserRepository(openTestDB(t))

	user := generateValidUser(t)
	err := repo.Create(context.Background(), user)
	require.NoError(t, err)

	t.Run("update existing user", func(t *testing.T) {
		user.Name = "Updated Name"
		user.Email = fmt.Sprintf("%s@mail.com", generateString(7))
		user.LastModified = time.Now().UTC()

		err := repo.Update(context.Background(), user)
		require.NoError(t, err)

		got, err := repo.Get(context.Background(), user.Id)
		require.NoError(t, err)
		require.Equal(t, user.Name, got.Name)
		require.Equal(t, user.Email, got.Email)
		require.WithinDuration(t, user.LastModified, got.LastModified, time.Second)
	})

	t.Run("update non-existing user", func(t *testing.T) {
		nonExistent := generateValidUser(t)
		err := repo.Update(context.Background(), nonExistent)
		require.Error(t, err)
	})

}

func TestUsersRepository_Delete(t *testing.T) {
  repo := postgres.NewUserRepository(openTestDB(t))

    t.Run("delete existing user", func(t *testing.T) {
        user := generateValidUser(t)
        err := repo.Create(context.Background(), user)
        require.NoError(t, err)

        err = repo.Delete(context.Background(), user.Id)
        require.NoError(t, err)

        got, err := repo.Get(context.Background(), user.Id)
        require.Error(t, err)
        require.Nil(t, got)
    })

    t.Run("delete non-existing user", func(t *testing.T) {
        nonExistentId := uuid.NewString()
        err := repo.Delete(context.Background(), nonExistentId)
        require.Error(t, err)
    })
}

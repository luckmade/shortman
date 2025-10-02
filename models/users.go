package models

import (
	"context"
	"time"
)

type User struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Password     string    `json:"-"`
	CreatedAt    time.Time `json:"createdAt"`
	LastModified time.Time `json:"lastModified"`
}

type UsersRepository interface {
	Create(ctx context.Context, user *User)error
	Get(ctx context.Context, id string)(*User, error)
	Update(ctx context.Context, user *User)error
	Delete(ctx context.Context, id string)error
}

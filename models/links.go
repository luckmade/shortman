package models

import (
	"context"
	"time"
)

type Link struct {
	Id           string    `json:"id"`
	UserId       string    `json:"userId"`
	LongURL      string    `json:"longUrl"`
	ShortURL     string    `json:"shortUrl"`
	CreatedAt    time.Time `json:"createdAt"`
	LastModified time.Time `json:"lastModified"`
}

type LinksRepository interface {
	Create(ctx context.Context, link *Link) error
	Get(ctx context.Context, id string) (*Link, error)
	Update(ctx context.Context, link *Link) error
	Delete(ctx context.Context, id string) error
}

package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jvhab/Redis-crud/internal/model"
	"github.com/redis/go-redis/v9"
)

type Repository struct {
	redisClient *redis.Client
}

func NewRepository(redisClient *redis.Client) *Repository {
	return &Repository{
		redisClient: redisClient,
	}
}

func (r *Repository) Create(ctx context.Context, news *model.News) (uuid.UUID, error) {

	return uuid.New(), nil
}

func (r *Repository) Update(ctx context.Context, news *model.News) error {
	return nil
}

func (r *Repository) Get(ctx context.Context, id uuid.UUID) (*model.News, error) {
	return nil, nil
}

func (r *Repository) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}

package repository

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/jvhab/Redis-crud/internal/model"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"time"
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
	id, err := uuid.NewUUID()
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "Repository.Create.uuid.NewUUID")
	}
	news.Id = id
	newsBytes, err := json.Marshal(news)
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "Repository.Create.json.Marshal")
	}
	if err := r.redisClient.Set(ctx, id.String(), newsBytes, time.Second*3).Err(); err != nil {
		return uuid.Nil, errors.Wrap(err, "Repository.Create.redis.Set")
	}
	return id, nil
}

func (r *Repository) Update(ctx context.Context, news *model.News) error {
	newBytes, err := json.Marshal(news)
	if err != nil {
		return errors.Wrap(err, "Repository.Update.json.Marshal")
	}
	if err := r.redisClient.Set(ctx, news.Id.String(), newBytes, time.Second*3).Err(); err != nil {
		return errors.Wrap(err, "Repository.Update.redis.Set")
	}

	return nil
}

func (r *Repository) Get(ctx context.Context, id uuid.UUID) (*model.News, error) {
	result := &model.News{}
	res, err := r.redisClient.Get(ctx, id.String()).Bytes()
	if err != nil {
		return &model.News{}, errors.Wrap(err, "Repository.Get.redis.Get")
	}
	if err := json.Unmarshal(res, result); err != nil {
		return &model.News{}, errors.Wrap(err, "Repository.Get.json.Unmarshal")
	}

	return result, nil
}

func (r *Repository) Delete(ctx context.Context, id uuid.UUID) error {
	if err := r.redisClient.Del(ctx, id.String()).Err(); err != nil {
		return errors.Wrap(err, "Repository.Del.redis.Del")
	}

	return nil
}

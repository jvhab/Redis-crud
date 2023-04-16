package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jvhab/Redis-crud/internal/model"
)

type RedisRepository interface {
	Create(context.Context, *model.News) (uuid.UUID, error)
	Update(context.Context, *model.News) error
	Get(context.Context, uuid.UUID) (*model.News, error)
	Delete(context.Context, uuid.UUID) error
}

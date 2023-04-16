package controller

import (
	"context"
	"github.com/google/uuid"
	"github.com/jvhab/Redis-crud/internal/model"
	"github.com/jvhab/Redis-crud/internal/repository"
)

type Controller struct {
	rep repository.RedisRepository
}

func NewController(ctrl repository.RedisRepository) *Controller {
	return &Controller{
		rep: ctrl,
	}
}

func (c *Controller) Create(ctx context.Context, news *model.News) error {
	return c.rep.Create(ctx, news)
}

func (c *Controller) Update(ctx context.Context, news *model.News) error {
	return c.rep.Update(ctx, news)
}

func (c *Controller) Get(ctx context.Context, id uuid.UUID) (*model.News, error) {
	return c.rep.Get(ctx, id)
}

func (c *Controller) Delete(ctx context.Context, id uuid.UUID) error {
	return c.rep.Delete(ctx, id)
}

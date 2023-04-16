package controller

import "github.com/jvhab/Redis-crud/internal/repository"

type Controller struct {
	rep repository.RedisRepository
}

func NewController(ctrl repository.RedisRepository) *Controller {
	return &Controller{
		rep: ctrl,
	}
}

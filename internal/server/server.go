package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jvhab/Redis-crud/config"
	"github.com/jvhab/Redis-crud/internal/controller"
	"github.com/jvhab/Redis-crud/internal/handler"
	"github.com/jvhab/Redis-crud/internal/repository"
	"github.com/jvhab/Redis-crud/pkg/redis"
	"log"
)

func Run(cfg *config.Config) {
	clientRedis := redis.NewRedisClient(cfg)
	repo := repository.NewRepository(clientRedis)
	ctrl := controller.NewController(repo)
	router := gin.New()
	hdl := handler.NewHandler(ctrl, router)
	hdl.RegisterRouter()

	err := router.Run(cfg.Server.Host)
	if err != nil {
		log.Fatal(err)
	}
}

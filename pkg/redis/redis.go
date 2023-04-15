package redis

import (
	"github.com/jvhab/Redis-crud/config"
	"github.com/redis/go-redis/v9"
	"time"
)

func NewRedisClient(cfg *config.Config) *redis.Client {
	option := &redis.Options{
		Addr:         cfg.Redis.RedisAddr,
		MinIdleConns: cfg.Redis.MinIdleConn,
		PoolSize:     cfg.Redis.PoolSize,
		PoolTimeout:  time.Duration(cfg.Redis.PoolTimeout) * time.Second,
		Password:     cfg.Redis.Password,
		DB:           cfg.Redis.DB,
	}
	client := redis.NewClient(option)

	return client
}
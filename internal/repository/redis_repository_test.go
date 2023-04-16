package repository

import (
	"context"
	"github.com/alicebob/miniredis/v2"
	"github.com/google/uuid"
	"github.com/jvhab/Redis-crud/internal/model"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func createMiniRedis() *Repository {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatal(err)
	}
	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})
	repo := NewRepository(client)

	return repo
}

func TestRedisRepository(t *testing.T) {
	t.Run("Test create news", func(t *testing.T) {
		repo := createMiniRedis()
		testNews := &model.News{
			AuthorId:  uuid.New(),
			Title:     "TestCreate",
			Content:   "Hard",
			Category:  "News",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		_, err := repo.Create(context.Background(), testNews)
		assert.NoError(t, err)
	})

	t.Run("Test Update news", func(t *testing.T) {
		repo := createMiniRedis()
		id, err := uuid.NewUUID()
		assert.NoError(t, err)
		testNews := &model.News{
			Id:        id,
			AuthorId:  uuid.New(),
			Title:     "TestUpdate",
			Content:   "Hard",
			Category:  "News",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		err = repo.Update(context.Background(), testNews)
		assert.NoError(t, err)
	})

	t.Run("Test Get news", func(t *testing.T) {
		repo := createMiniRedis()
		testNews := &model.News{
			AuthorId:  uuid.New(),
			Title:     "TestGet",
			Content:   "Hard",
			Category:  "News",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		id, err := repo.Create(context.Background(), testNews)
		assert.NoError(t, err)
		result, err := repo.Get(context.Background(), id)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("Test Delete news", func(t *testing.T) {
		repo := createMiniRedis()
		testNews := &model.News{
			AuthorId:  uuid.New(),
			Title:     "TestDel",
			Content:   "Hard",
			Category:  "News",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		id, err := repo.Create(context.Background(), testNews)
		assert.NoError(t, err)
		err = repo.Delete(context.Background(), id)
		assert.NoError(t, err)
	})
}

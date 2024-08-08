package repository

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisRepositoryInterface interface {
	GetRepoData(ctx context.Context, key string) (string, error)
	PostRepoData(ctx context.Context, key string, value any) error
}

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(ep string) (*RedisRepository, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr: ep,
	})
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return &RedisRepository{client: redisClient}, nil
}

func (r *RedisRepository) GetRepoData(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *RedisRepository) PostRepoData(ctx context.Context, key string, value any) error {
	res := r.client.Set(ctx, key, value, 30*time.Second)
	if res.Err() != nil {
		return res.Err()
	}
	return nil
}

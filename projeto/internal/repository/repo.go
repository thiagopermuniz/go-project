package repository

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisClientAPI interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
}

type RedisRepositoryInterface interface {
	GetRepoData(ctx context.Context, key string) (string, error)
	SetRepoData(ctx context.Context, key, value string) (string, error)
}

type RedisRepository struct {
	client RedisClientAPI
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

func (r *RedisRepository) SetRepoData(ctx context.Context, key, value string) (string, error) {
	res := r.client.Set(ctx, key, value, 30*time.Second)
	if res.Err() != nil {
		return "", res.Err()
	}
	return res.Val(), nil
}

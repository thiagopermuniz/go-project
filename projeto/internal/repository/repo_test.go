package repository

import (
	"context"
	"github.com/redis/go-redis/v9"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRedisClient struct {
	mock.Mock
}

func (m *MockRedisClient) Get(ctx context.Context, key string) *redis.StringCmd {
	args := m.Called(ctx, key)
	return args.Get(0).(*redis.StringCmd)
}

func (m *MockRedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	args := m.Called(ctx, key, value, expiration)
	return args.Get(0).(*redis.StatusCmd)
}

func TestRedisRepository_GetData(t *testing.T) {
	mockClient := new(MockRedisClient)
	repo := &RedisRepository{client: mockClient}

	ctx := context.Background()
	key := "test-key"
	expectedValue := "test-value"

	stringCmd := redis.NewStringCmd(ctx, key)
	stringCmd.SetVal(expectedValue)

	mockClient.On("Get", ctx, key).Return(stringCmd)

	value, err := repo.GetRepoData(ctx, key)

	assert.NoError(t, err)
	assert.Equal(t, expectedValue, value)

	mockClient.AssertExpectations(t)
}

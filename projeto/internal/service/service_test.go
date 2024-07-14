package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRedisRepository struct {
	mock.Mock
}

func (m *MockRedisRepository) GetRepoData(ctx context.Context, key string) (string, error) {
	args := m.Called(ctx, key)
	return args.String(0), args.Error(1)
}

func (m *MockRedisRepository) SetRepoData(ctx context.Context, key, value string) (string, error) {
	args := m.Called(ctx, key, value)
	return args.String(0), args.Error(1)
}

func TestDataService_GetData(t *testing.T) {
	mockRepo := new(MockRedisRepository)
	dataService := NewDataService(mockRepo)

	ctx := context.Background()
	key := "test-key"
	expectedValue := "test-value"

	mockRepo.On("GetRepoData", ctx, key).Return(expectedValue, nil)

	value, err := dataService.GetServiceData(ctx, key)

	assert.NoError(t, err)
	assert.Equal(t, expectedValue, value)

	mockRepo.AssertExpectations(t)
}

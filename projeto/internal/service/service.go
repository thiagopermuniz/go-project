package service

import (
	"context"
	"projeto/internal/repository"
)

type DataServiceInterface interface {
	GetServiceData(ctx context.Context, id string) (string, error)
	SetServiceData(ctx context.Context, key, value string) (string, error)
}

func NewDataService(repo repository.RedisRepositoryInterface) *DataService {
	return &DataService{
		repository: repo,
	}
}

type DataService struct {
	repository repository.RedisRepositoryInterface
}

func (s *DataService) GetServiceData(ctx context.Context, key string) (string, error) {
	return s.repository.GetRepoData(ctx, key)
}

func (s *DataService) SetServiceData(ctx context.Context, key, value string) (string, error) {
	return s.repository.SetRepoData(ctx, key, value)
}

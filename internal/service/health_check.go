package service

import (
	"context"

	"insightGlobal_carInventory/internal/repository"
)

// HealthCheckService specific methods
type HealthCheckService interface {
	PingMysql(ctx context.Context) error
}

type healthCheckService struct {
	repository repository.Repository
}

func newHealthCheckService(r repository.Repository) *healthCheckService {
	return &healthCheckService{repository: r}
}

// PingMysql pings database using the connection to check if it's working properly
func (s healthCheckService) PingMysql(ctx context.Context) error {
	return s.repository.Ping(ctx)
}

package service

import (
	"context"

	"transaction-service/internal/repository"
)

// HealthCheckService specific methods
type HealthCheckService interface {
	PingMysql(ctx context.Context) error
}

type healthCheckService struct {
	repository repository.Repositories
}

func newHealthCheckService(r repository.Repositories) *healthCheckService {
	return &healthCheckService{repository: r}
}

// PingMysql pings database using the connection to check if it's working properly
func (s healthCheckService) PingMysql(ctx context.Context) error {
	return s.repository.Ping(ctx)
}

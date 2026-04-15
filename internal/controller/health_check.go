package controller

import (
	"context"

	"transaction-service/internal/service"
)

// HealthCheckController specific methods
type HealthCheckController interface {
	PingDB(context.Context) error
}

type healthCheckController struct {
	healthCheckService service.UserTransaction
}

func newHealthCheckController(s service.UserTransaction) *healthCheckController {
	return &healthCheckController{s}
}

func (c healthCheckController) PingDB(ctx context.Context) error {
	return c.healthCheckService.HealthCheckService(ctx)
}

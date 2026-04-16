package handler

import (
	"context"
	"insightGlobal_carInventory/internal/service"

	"go.uber.org/zap"
)

type HealthCheckHandler interface {
	CheckServices(ctx context.Context) error
}

type healthCheckHandler struct {
	healthCheckService service.HealthCheckService
	log                *zap.SugaredLogger
}

func newHealthCheckHandler(service service.HealthCheckService, log *zap.SugaredLogger) HealthCheckHandler {
	return &healthCheckHandler{
		healthCheckService: service,
		log:                log,
	}
}

func (c healthCheckHandler) CheckServices(ctx context.Context) error {
	return c.healthCheckService.PingMysql(ctx)
}

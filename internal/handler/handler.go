package handler

import (
	"insightGlobal_carInventory/internal/infrastructure"
	"insightGlobal_carInventory/internal/repository"
	"insightGlobal_carInventory/internal/service"

	"go.uber.org/zap"
)

// Handlers base methods
type Handlers interface {
	HealthCheckHandler() HealthCheckHandler
	CarHandler() CarHandler
}

type handlers struct {
	healthCheckHandler HealthCheckHandler
	carHandler         CarHandler
}

// NewHandlers returns Handlers with all required dependencies
func NewHandlers(dbConn infrastructure.DB, log *zap.SugaredLogger) Handlers {
	repositories := repository.NewRepositories(dbConn, log)

	return &handlers{
		healthCheckHandler: newHealthCheckHandler(service.NewHealthCheckService(repositories), log),
		carHandler:         newCarHandler(service.NewCarService(repositories, log), log),
	}
}

func (h *handlers) CarHandler() CarHandler {
	return h.carHandler
}

func (h *handlers) HealthCheckHandler() HealthCheckHandler {
	return h.healthCheckHandler
}

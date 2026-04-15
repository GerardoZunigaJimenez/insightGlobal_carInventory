package controller

import (
	"go.uber.org/zap"
)

// Controllers base methods
type Controllers interface {
	HealthCheckController() HealthCheckController
	CarController() CarController
}

type controller struct {
	healthCheckController HealthCheckController
	carController         CarController
}

// NewControllers returns controllers with all required dependencies
func NewControllers(l *zap.SugaredLogger) Controllers {
	return &controller{
		healthCheckController: newHealthCheckController(sut),
		carController:         newCarController(sut, l),
	}
}

func (c controller) CarController() CarController {
	return c.carController
}

func (c controller) HealthCheckController() HealthCheckController {
	return c.healthCheckController
}

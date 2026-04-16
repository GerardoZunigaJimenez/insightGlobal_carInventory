package app

import (
	"net/http"

	bun "github.com/uptrace/bunrouter"
)

const (
	healthcheck                           = "/healthcheck"
	CarPathGroup                          = "/cars"
	ContentTypeHeaderKey                  = "Content-Type"
	ContentTypeApplicationJsonHeaderValue = "application/json"
)

func (a *app) setHealthRoute() {

	a.router.GET(healthcheck, func(w http.ResponseWriter, req bun.Request) error {
		err := a.handlers.HealthCheckHandler().CheckServices(req.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return bun.JSON(w, bun.H{"message": "error consuming healthcheck services"})
		}
		return bun.JSON(w, bun.H{"status": "Alive All Services"})
	})
}

func (a *app) setAPIRoutes() {
	// Transaction group middleware
	carGroup := a.router.NewGroup(CarPathGroup).Use(func(next bun.HandlerFunc) bun.HandlerFunc {
		return func(w http.ResponseWriter, req bun.Request) error {
			w.Header().Add(ContentTypeHeaderKey, ContentTypeApplicationJsonHeaderValue)
			return next(w, req)
		}
	})
	// POST /cars - Create a new car
	carGroup.POST("", a.handlers.CarHandler().Create)
	// GET /cars/:id - Get car by ID
	carGroup.GET("/:id", a.handlers.CarHandler().GetByID)
	// GET /cars - List cars with pagination
	carGroup.GET("", a.handlers.CarHandler().List)
	// PUT /cars/:id - Update a car
	carGroup.PUT("/:id", a.handlers.CarHandler().Update)
	// DELETE /cars/:id - Delete a car
	carGroup.DELETE("/:id", a.handlers.CarHandler().Delete)
}

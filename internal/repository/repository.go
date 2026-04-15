package repository

import (
	"context"
	"insightGlobal_carInventory/internal/infrastructure"

	"go.uber.org/zap"
)

//go:generate mockery
type Repository interface {
	Ping(context.Context) error
	CarStorage() CarStorage
}

type repositories struct {
	carStorage CarStorage
	ping       Ping
}

func NewRepositories(dbConn infrastructure.DB, log *zap.SugaredLogger) Repository {
	return &repositories{
		ping:       newPing(dbConn),
		carStorage: newCarStorage(dbConn, log),
	}
}

func (r *repositories) Ping(ctx context.Context) error {
	return r.ping.Ping(ctx)
}

func (r *repositories) CarStorage() CarStorage {
	return r.carStorage
}

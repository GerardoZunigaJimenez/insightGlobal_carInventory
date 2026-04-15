package repository

import (
	"context"
	"insightGlobal_carInventory/internal/infrastructure"

	"go.uber.org/zap"
)

// Repositories base methods
type Repositories interface {
	CarStorage() CarStorage
}

type repository struct {
	car CarStorage
}

func NewRepositories(dbConn *infrastructure.DB, log *zap.SugaredLogger) Repositories {
	return &repository{
		carStorage: newCarStorage(dbConn, log),
	}
}

func (r repository) Ping(ctx context.Context) error {
	if err := r.Ping(); err != nil {
		return err
	}

	return nil
}

package repository

import (
	"context"
	"fmt"
	"insightGlobal_carInventory/internal/infrastructure"
)

type Ping interface {
	Ping(context.Context) error
}

type ping struct {
	db infrastructure.DB
}

func (p ping) Ping(ctx context.Context) error {
	if err := p.db.IsAlive(); err != nil {
		return fmt.Errorf("unable to ping the DB, because the reference is nil")
	}
	return p.db.PingContext()
}

func newPing(dbConn infrastructure.DB) *ping {
	return &ping{dbConn}
}

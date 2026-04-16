package infrastructure

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"insightGlobal_carInventory/internal/config"

	_ "github.com/lib/pq" // Side-effect import to register the driver
	gitErrors "github.com/pkg/errors"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"go.uber.org/zap"
)

const (
	postgresDriverName = "postgres"
	dbConnectionError  = "failed to set up database client"
	dbNotReadyError    = "db connection has not been set up"
	dbProviderName     = "Postgres Provider"
)

// DB is an adapter for the database ORM.
type DB interface {
	PingContext() error
	IsAlive() error
	Connect() error
	Select() *bun.SelectQuery
	Insert() *bun.InsertQuery
	Update() *bun.UpdateQuery
}

type db struct {
	ctx    context.Context
	client *bun.DB
	conn   *bun.Conn
}

// NewBunPostgresClient returns a new bun connection and db client.
func NewBunPostgresClient(ctx context.Context, log *zap.SugaredLogger, dbConfig *config.DB) DB {
	log.Info("Setting DB connection ...")

	dsn := buildPostgresDSN(dbConfig)
	sqlClient, err := sql.Open(postgresDriverName, dsn)
	if err != nil {
		log.Fatal(&sql.DB{}, fmt.Errorf("%s: %w", dbConnectionError, err))
	}

	database := &db{
		ctx:    ctx,
		client: bun.NewDB(sqlClient, pgdialect.New()),
	}

	if err := database.Connect(); err != nil {
		log.Fatal("failed to initialize database connection:", err)
	}

	return database
}

func buildPostgresDSN(dbConfig *config.DB) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Name,
	)
}

func (d *db) Connect() error {
	conn, err := d.client.Conn(d.ctx)
	if err != nil {
		return gitErrors.Wrap(err, dbProviderName)
	}

	d.conn = &conn
	return nil
}

func (d *db) IsAlive() error {
	if d.conn == nil {
		return errors.New(dbNotReadyError)
	}
	return nil
}

func (d *db) PingContext() error {
	return d.client.PingContext(d.ctx)
}

func (d *db) Select() *bun.SelectQuery {
	return d.client.NewSelect()
}

func (d *db) Insert() *bun.InsertQuery {
	return d.client.NewInsert()
}

func (d *db) Update() *bun.UpdateQuery {
	return d.client.NewUpdate()
}

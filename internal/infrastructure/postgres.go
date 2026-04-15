package infrastructure

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	gitErrors "github.com/pkg/errors"

	"insightGlobal_carInventory/internal/config"

	_ "github.com/lib/pq" // Side-effect import to register the driver
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"go.uber.org/zap"
)

// DB is an adapter for the database ORM
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
	log    *zap.SugaredLogger
}

// NewBunMysqlClient returns a new bun connection and db Client (connection is used for pinging)
func NewBunMysqlClient(ctx context.Context, log *zap.SugaredLogger, config *config.DB) DB {
	log.Info("Setting DB connection ...")
	var database db
	database.ctx = ctx
	database.log = log
	sqlClient, err := sql.Open("postgres", getMysqlDSN(config))
	if err != nil {
		log.Fatal(&sql.DB{}, fmt.Errorf("failed to set up Client %v", err))
	}
	database.client = bun.NewDB(sqlClient, mysqldialect.New())

	return &database
}

func getMysqlDSN(config *config.DB) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.User, config.Password, config.Host, config.Port, config.Name)
}

func (d *db) Connect() error {
	conn, err := d.client.Conn(d.ctx)

	if err != nil {
		return gitErrors.Wrap(err, "MySQL Provider")
	}

	d.conn = &conn

	return nil
}

func (d *db) IsAlive() error {
	if d.conn == nil {
		return errors.New("db connection has not been set up")
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

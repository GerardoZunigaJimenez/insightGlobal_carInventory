package app

import (
	"context"
	"errors"
	"fmt"
	"insightGlobal_carInventory/internal/handler"
	"insightGlobal_carInventory/internal/infrastructure"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"insightGlobal_carInventory/internal/config"

	bun "github.com/uptrace/bunrouter"
	"go.uber.org/zap"
	// necessary Postgres driver
	_ "github.com/lib/pq"
)

const (
	ServerShutdownTimeout = 5 * time.Second
)

type app struct {
	config   *config.Application
	handlers handler.Handlers
	log      *zap.SugaredLogger
	ctx      context.Context
	router   *bun.Router
}

func (a *app) start() {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", a.config.ServerPort),
		Handler: a.router,
	}

	go func() {
		// service connections
		a.log.Info(fmt.Sprintf("Listening %s...", a.config.ServerPort))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	a.log.Info("Server started successfully!")
	<-quit
	a.log.Info("Shutting down the server ...")

	ctx, cancel := context.WithTimeout(a.ctx, 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		a.log.Fatal("Server Shutdown:", err)
	}

	a.log.Info(fmt.Sprintf("(%.0f) second(s) to complete the shutdown...", ServerShutdownTimeout.Seconds()))
	<-ctx.Done()
	a.log.Info("terminated!")
}

// ServeAPI creates api, sets up routes, runs HTTP server and performs necessary clean tasks afterwards
func ServeAPI() {
	ctx := context.Background()

	// log set up
	log := zap.Must(zap.NewDevelopment()).Sugar()

	appConfig := config.NewApplication()
	if appConfig == nil {
		log.Fatal("Application configuration is nil, cannot proceed with server setup")
	}

	if !strings.EqualFold(appConfig.Environment, "local") {
		infrastructure.NewGCPConnection(ctx, config.NewGCP())
	}
	dbConfig := config.NewDB()
	dbConn := infrastructure.NewBunPostgresClient(ctx, log, dbConfig)
	handlers := handler.NewHandlers(dbConn, log)

	var app = &app{
		config:   appConfig,
		handlers: handlers,
		ctx:      ctx,
		log:      log,
		router:   bun.New(bun.Use()),
	}

	app.setHealthRoute()
	app.setAPIRoutes()

	app.start()
}

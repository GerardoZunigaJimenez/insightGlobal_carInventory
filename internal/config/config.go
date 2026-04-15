package config

import (
	"github.com/kelseyhightower/envconfig"
)

type ApiEndpoint string

// Application configuration environment variables used for logging, tracing and overall monitoring
type Application struct {
	Name        string `envconfig:"NAME"`
	Version     string `envconfig:"VERSION"`
	Environment string `envconfig:"ENVIRONMENT"`
	ServerPort  string `envconfig:"SERVER_PORT"`
}

// GCP configuration environment variables
type GCP struct {
	ProjectID              string `envconfig:"PROJECT_ID"`
	InstanceConnectionName string `envconfig:"INSTANCE_CONNECTION_NAME"`
}

// DB Cloud DB ir Local Postgress DB
type DB struct {
	Host     string `envconfig:"HOST"`
	Port     string `envconfig:"PORT"`
	User     string `envconfig:"USER"`
	Password string `envconfig:"PASSWORD"`
	Name     string `envconfig:"NAME"`
}

// NewApplication return a new application config by reading the environment variables
func NewApplication() *Application {
	config := new(Application)
	envconfig.MustProcess("app", config)
	return config
}

// NewGCP New GCP returns a new AWS config by reading the environment variables
func NewGCP() *GCP {
	config := new(GCP)
	envconfig.MustProcess("gcp", config)
	return config
}

// NewDB returns a database config by reading the environment variables
func NewDB() *DB {
	config := new(DB)
	envconfig.MustProcess("db", config)
	return config
}

package config

import (
	"time"
)

type AppConfig struct {
	AppPort                       int           `env:"APP_PORT" envDefault:"8000"`
	GracefulShutdownPeriodSeconds time.Duration `env:"GRACEFUL_SHUTDOWN_PERIOD_SECONDS" envDefault:"5s"`
}

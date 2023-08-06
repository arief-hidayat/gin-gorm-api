package main

import (
	"context"
	"fmt"
	"github.com/arief-hidayat/gin-gorm-api/adapter"
	"github.com/arief-hidayat/gin-gorm-api/config"
	"github.com/arief-hidayat/gin-gorm-api/internal/route"
	"github.com/arief-hidayat/gin-gorm-api/pkg/logger"
	"github.com/caarlos0/env/v9"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"os/signal"
	"syscall"
)

var writerDb, readerDb *gorm.DB = adapter.ConnectWithDB()

// @title           Gin GORM Sample API
// @version         1.0
// @description     Gin GORM Sample API

// @contact.name   Arief Hidayat
// @contact.url    https://linkedin.com/in/ariefh
// @contact.email  mr.arief.hidayat@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /api/v1

func main() {
	defer adapter.CloseDbConnection(writerDb)
	defer adapter.CloseDbConnection(readerDb)
	logr := logger.NewLogger()
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := gin.Default()
	route.RootRoute(writerDb, readerDb, router, logr)
	appConfig := config.AppConfig{}
	if err := env.Parse(&appConfig); err != nil {
		fmt.Printf("%+v\n", err)
	}
	logr.Info().Msg(fmt.Sprintf("Running on :%d", appConfig.AppPort))
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", appConfig.AppPort),
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logr.Error().Err(err).Msg("listen")
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	logr.Info().Msg("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), appConfig.GracefulShutdownPeriodSeconds)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logr.Error().Err(err).Msg("Server forced to shutdown: ")
	}

	logr.Info().Msg("Server exiting")
}

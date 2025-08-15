package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/weeweeshka/tataisk/internal/app/buildApp"
	"github.com/weeweeshka/tataisk/internal/config"
	"github.com/weeweeshka/tataisk/pkg/lib/logger"
	"go.uber.org/zap"
)

func main() {
	cfg := config.MustLoadConfig()

	logr := logger.SetupLogger()

	logr.Info("Config loaded")
	logr.Info("Logger initialized")
	app, err := buildApp.NewApp(cfg.Port, cfg.StoragePath, logr)
	if err != nil {
		logr.Info("failed to initialize app")
		panic(err)
	}

	go func() {
		if err := app.GRPCServer.MustRun(); err != nil {
			logr.Info("failed to run server", zap.Error(err))
		}
	}()
	logr.Info("Server started", zap.String("port", cfg.Port))

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	slog.Info("App stopped")
	app.GRPCServer.GracefulStop()
}

package main

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/weeweeshka/tataisk/internal/config"
	"github.com/weeweeshka/tataisk/pkg/lib/logger"
)

func main() {
	cfg := config.MustLoadConfig()
	logr := logger.SetupLogger()

	m, err := migrate.New("tataisk/migrations", cfg.StoragePath)
	if err != nil {
		logr.Info("Migrations initialisation failure")
		panic(err)
	}

	if err := m.Up(); err != nil {
		logr.Info("Migrations not applied")
	}
}

package buildApp

import (
	"fmt"
	runGrpc "github.com/weeweeshka/tataisk/internal/app/grpcApp"
	"github.com/weeweeshka/tataisk/internal/repository/postgres"
	service "github.com/weeweeshka/tataisk/internal/services/tataisk"
	"go.uber.org/zap"
)

type App struct {
	GRPCServer runGrpc.GRPCServer
}

func NewApp(port string, storagePath string, logr *zap.Logger) (*App, error) {
	storage, err := postgres.NewStorage(storagePath, logr)
	if err != nil {
		return &App{}, fmt.Errorf("failed to run db : %w", err)
	}
	tataiskService := service.New(logr, storage)
	grpcApp := runGrpc.NewGRPCServer(port, logr, tataiskService)

	return &App{GRPCServer: *grpcApp}, nil
}

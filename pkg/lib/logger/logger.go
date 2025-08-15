package logger

import (
	"go.uber.org/zap"
)

func SetupLogger() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic("failed to initialize logger %v")
	}

	return logger
}

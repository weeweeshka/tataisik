package grpcApp

import (
	"fmt"
	grpcHandlers "github.com/weeweeshka/tataisk/internal/grpcHandlers"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type GRPCServer struct {
	gRPCServer *grpc.Server
	port       string
	logr       *zap.Logger
}

func NewGRPCServer(port string, logr *zap.Logger, tataiskService grpcHandlers.Tataisk) *GRPCServer {
	grpcServer := grpc.NewServer()
	grpcHandlers.RegisterNewServer(grpcServer, tataiskService)
	return &GRPCServer{
		gRPCServer: grpcServer,
		port:       port,
		logr:       logr,
	}

}

func (s *GRPCServer) MustRun() error {
	l, err := net.Listen("tcp", ":"+s.port)
	if err != nil {
		return fmt.Errorf("cannot listen: %w", err)
	}
	err = s.gRPCServer.Serve(l)
	if err != nil {
		return fmt.Errorf("cannot serve: %w", err)
	}

	s.logr.Info("grpc server started", zap.String("port", s.port))
	return nil
}

func (s *GRPCServer) GracefulStop() {
	s.gRPCServer.GracefulStop()
}

package main

import (
	"context"
	"exchange-rate-receiver/internal/config"
	"exchange-rate-receiver/internal/grpc"
	"exchange-rate-receiver/internal/repository"
	"exchange-rate-receiver/internal/service"
	"github.com/jackc/pgx/v4/pgxpool"
	"net"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("Failed to load config", zap.Error(err))
	}

	dbPool, err := pgxpool.Connect(context.Background(), cfg.DatabaseURL)
	if err != nil {
		logger.Fatal("Unable to connect to database", zap.Error(err))
	}
	defer dbPool.Close()

	repo := repository.NewRepository(dbPool)
	srv := service.NewService(repo)
	grpcServer := grpc.NewServer(srv, logger)
	lis, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		logger.Fatal("Failed to listen", zap.Error(err))
	}

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			logger.Fatal("Failed to serve GRPC server", zap.Error(err))
		}
	}()

	// Graceful shutdown
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	grpcServer.Stop()
}

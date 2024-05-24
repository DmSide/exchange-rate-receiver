package grpc

import (
	"context"
	"exchange-rate-receiver/internal/service"
	"net"

	_ "exchange-rate-receiver/internal/service"
	pb "exchange-rate-receiver/proto"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedExchangeRateServiceServer
	svc        service.Service
	grpcServer *grpc.Server
	logger     *zap.Logger
}

func NewServer(svc service.Service, logger *zap.Logger) *Server {
	return &Server{
		svc:        svc,
		grpcServer: grpc.NewServer(),
		logger:     logger,
	}
}

func (s *Server) Serve(lis net.Listener) error {
	pb.RegisterExchangeRateServiceServer(s.grpcServer, s)
	s.logger.Info("GRPC server listening", zap.String("address", lis.Addr().String()))
	return s.grpcServer.Serve(lis)
}

func (s *Server) Stop() {
	s.logger.Info("Stopping GRPC server")
	s.grpcServer.GracefulStop()
}

func (s *Server) GetRates(ctx context.Context, req *pb.GetRatesRequest) (*pb.GetRatesResponse, error) {
	rates, err := s.svc.GetRates()
	if err != nil {
		s.logger.Error("Failed to get rates", zap.Error(err))
		return nil, err
	}
	return &pb.GetRatesResponse{
		Ask:       rates.Ask,
		Bid:       rates.Bid,
		Timestamp: rates.Timestamp,
	}, nil
}

func (s *Server) Healthcheck(ctx context.Context, req *pb.HealthcheckRequest) (*pb.HealthcheckResponse, error) {
	s.logger.Info("Healthcheck called")
	return &pb.HealthcheckResponse{Status: "Healthy"}, nil
}

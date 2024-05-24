package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()

	// Graceful shutdown
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	log.Println("Stopping GRPC server...")
	grpcServer.GracefulStop()
	log.Println("Server stopped.")
}

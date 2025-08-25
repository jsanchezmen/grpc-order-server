package grpc

import (
	"fmt"
	"grpc-order-server/config"
	go_src "grpc-order-server/internal/adapters/grpc/proto/go-src"
	"log"
	"log/slog"
	"net"

	"google.golang.org/grpc/reflection"
)

func (a Adapter) Run() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("Failed to listen on port %d error: %v", a.port, err)
	}
	go_src.RegisterOrderServer(a.server, a)

	if config.GetEnv() == "dev" {
		reflection.Register(a.server)
	}
	slog.Info("Starting gRPC server", "port", a.port, "ENV", config.GetEnv())

	if err := a.server.Serve(listen); err != nil {
		log.Fatalf("Failed to serve grpc on port %v", a.port)
	}

}

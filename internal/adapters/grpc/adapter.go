package grpc

import (
	go_src "grpc-order-server/internal/adapters/grpc/proto/go-src"
	"grpc-order-server/internal/application/ports"

	"google.golang.org/grpc"
)

type Adapter struct {
	api    ports.Apiport
	port   int
	server *grpc.Server
	go_src.UnimplementedOrderServer
}

func NewAdapter(api ports.Apiport, port int) *Adapter {
	return &Adapter{
		api:    api,
		port:   port,
		server: grpc.NewServer(),
	}
}

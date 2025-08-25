package main

import (
	"grpc-order-server/config"
	"grpc-order-server/internal/adapters/grpc"
	"grpc-order-server/internal/application/api"
)

func main() {
	application := api.NewApplication()

	grpcAdapter := grpc.NewAdapter(application, config.GetServerPort())
	grpcAdapter.Run()

}

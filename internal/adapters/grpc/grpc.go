package grpc

import (
	"context"
	"fmt"
	go_src "grpc-order-server/internal/adapters/grpc/proto/go-src"
	"io"
	"log/slog"
	"time"

	"google.golang.org/grpc"
)

func (a Adapter) Create(ctx context.Context, request *go_src.CreateOrderRequest) (*go_src.CreateOrderResponse, error) {
	itemName := request.GetItemName()

	orderCreated, err := a.api.CreateOrder(itemName)
	if err != nil {
		return nil, err
	}

	return &go_src.CreateOrderResponse{
		OrderId:   orderCreated.Id,
		ItemName:  orderCreated.ItemName,
		CreatedAt: orderCreated.CreatedAt,
	}, nil
}

func (a Adapter) ListOrders(request *go_src.CreateListOrderRequest, stream grpc.ServerStreamingServer[go_src.CreateOrderResponse]) error {
	ordersQuantity := int(request.GetOrdersQuantity())

	for i := 0; i <= ordersQuantity; i++ {
		orderCreated, err := a.api.CreateOrder(fmt.Sprintf("Item %d", i))
		if err != nil {
			slog.Error("Failed to create order", "error", err)
			return err
		}
		stream.SendMsg(&go_src.CreateOrderResponse{
			OrderId:   int64(i),
			ItemName:  orderCreated.ItemName,
			CreatedAt: orderCreated.CreatedAt,
		})
		//simulate delay for streaming
		time.Sleep(1 * time.Second)

	}

	return nil
}

func (a Adapter) CreateStreamOrder(stream grpc.ClientStreamingServer[go_src.CreateOrderRequest, go_src.CreateStreamOrderResponse]) error {
	ordersCreated := []*go_src.CreateOrderResponse{}
	for {
		orderRequest, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&go_src.CreateStreamOrderResponse{
				Orders:             ordersCreated,
				TotalOrdersCreated: int64(len(ordersCreated)),
			})
		}
		if err != nil {
			slog.Error("Failed to receive order Stream request", "error", err)
			return err
		}
		itemName := orderRequest.GetItemName()
		orderCreated, err := a.api.CreateOrder(itemName)

		ordersCreated = append(ordersCreated, &go_src.CreateOrderResponse{
			OrderId:   int64(len(ordersCreated) + 1),
			ItemName:  orderCreated.ItemName,
			CreatedAt: orderCreated.CreatedAt,
		})

	}

}

func (a Adapter) CreateBidirectionalStreamOrder(stream grpc.BidiStreamingServer[go_src.CreateOrderRequest, go_src.CreateOrderResponse]) error {
	cont := 1
	for {
		orderRequest, err := stream.Recv()

		if err == io.EOF {
			return nil // End of stream
		}

		if err != nil {
			slog.Error("Failed to receive order request", "error", err)
			return err
		}
		itemName := orderRequest.GetItemName()
		orderCreated, err := a.api.CreateOrder(itemName)
		sendError := stream.Send(&go_src.CreateOrderResponse{
			OrderId:   int64(cont),
			ItemName:  orderCreated.ItemName,
			CreatedAt: orderCreated.CreatedAt,
		})
		if sendError != nil {
			slog.Error("Failed to send order response", "error", sendError)
			return sendError
		}
		//simulate delay for streaming
		time.Sleep(1 * time.Second)
		cont++
	}
}

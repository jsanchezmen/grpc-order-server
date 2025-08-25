package ports

import "grpc-order-server/internal/application/domain"

type Apiport interface {
	CreateOrder(itemName string) (domain.Order, error)
}

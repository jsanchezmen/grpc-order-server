package api

import "grpc-order-server/internal/application/domain"

type Application struct {
}

func NewApplication() *Application {
	return &Application{}
}

func (a *Application) CreateOrder(itemName string) (domain.Order, error) {
	order := domain.NewOrder(itemName)
	return order, nil
}

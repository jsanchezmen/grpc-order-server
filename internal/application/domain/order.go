package domain

import "time"

type Order struct {
	Id        int64  `json:"id"`
	ItemName  string `json:"item_name"`
	CreatedAt int64  `json:"created_at"`
}

func NewOrder(ItemName string) Order {
	return Order{
		CreatedAt: time.Now().Unix(),
		ItemName:  ItemName,
		Id:        1,
	}
}

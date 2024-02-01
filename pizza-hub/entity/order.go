package entity

import "time"

type Order struct {
	ID    uint64      `json:"id"`
	Items []OrderItem `json:"items"`
}
type OrderItem struct {
	MenuID     uint64    `json:"menu_id"`
	ChefID     uint64    `json:"chef_id"`
	CreatedAt  time.Time `json:"created_at"`
	FinishedAt time.Time `json:"finished_at"`
}
type OrderRepository interface {
	Create(order *Order) error
}

type OrderService interface {
	CreateOrder(order *Order) error
}

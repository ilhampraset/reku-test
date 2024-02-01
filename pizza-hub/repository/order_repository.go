package repository

import (
	"errors"
	"reku-code-test/pizza-hub/entity"
	"sync"
)

type MemoryOrderRepository struct {
	orders map[uint64]*entity.Order
	mu     sync.Mutex
}

func NewMemoryOrderRepository() *MemoryOrderRepository {
	return &MemoryOrderRepository{
		orders: make(map[uint64]*entity.Order),
	}
}

func (r *MemoryOrderRepository) Create(order *entity.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.orders[order.ID]; exists {
		return errors.New("order with the same ID already exists")
	}

	r.orders[order.ID] = order
	return nil
}

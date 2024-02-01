package repository

import (
	"reflect"
	"reku-code-test/pizza-hub/entity"
	"testing"
)

func TestMemoryOrderRepository_Create(t *testing.T) {
	repo := NewMemoryOrderRepository()

	order := &entity.Order{
		ID: 1,
		Items: []entity.OrderItem{
			{MenuID: 1},
			{MenuID: 2},
		},
	}

	err := repo.Create(order)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	err = repo.Create(order)
	if err == nil {
		t.Error("Expected an error, but got none")
	}

	storedOrder, exists := repo.orders[order.ID]
	if !exists {
		t.Error("Order not found in the repository")
	}

	if !reflect.DeepEqual(storedOrder, order) {
		t.Errorf("Expected order: %v, but got: %v", order, storedOrder)
	}
}

package service

import (
	"reku-code-test/pizza-hub/entity"
	"sync"
	"testing"
)

type MockOrderRepository struct {
	orders map[uint64]*entity.Order
}

func NewMockOrderRepository() *MockOrderRepository {
	return &MockOrderRepository{
		orders: make(map[uint64]*entity.Order),
	}
}

func (r *MockOrderRepository) Create(order *entity.Order) error {
	r.orders[order.ID] = order
	return nil
}

func (r *MockOrderRepository) All() []entity.Order {
	orders := make([]entity.Order, 0, len(r.orders))
	for _, order := range r.orders {
		orders = append(orders, *order)
	}
	return orders
}

type MockMenuRepository struct{}

func (r *MockMenuRepository) Add(menu *entity.Menu) error {

	return nil
}

func (r *MockMenuRepository) All() []entity.Menu {
	return []entity.Menu{
		{ID: 1, Name: "Pizza Cheese", CookingTime: 2},
		{ID: 2, Name: "Pizza Meatball", CookingTime: 3},
	}
}

type MockChefRepository struct{}

func (r *MockChefRepository) Add(chef *entity.Chef) error {
	return nil
}

func (r *MockChefRepository) All() ([]entity.Chef, error) {
	return []entity.Chef{{ID: 1, Name: "John"}, {ID: 2, Name: "Doe"}}, nil
}

func setup() *orderService {
	return NewOrderService(
		NewMockOrderRepository(),
		&MockMenuRepository{},
		&MockChefRepository{},
	)
}

func TestOrderService_CreateOrder(t *testing.T) {
	orderService := setup()

	mockOrder := &entity.Order{
		ID: 1,
		Items: []entity.OrderItem{
			{MenuID: 1},
			{MenuID: 2},
			{MenuID: 1},
		},
	}

	err := orderService.CreateOrder(mockOrder)
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}
}

func TestAssignChef(t *testing.T) {
	orderService := setup()

	chefs := []entity.Chef{{ID: 1, Name: "John"}, {ID: 2, Name: "Jane"}}
	order := &entity.Order{
		Items: []entity.OrderItem{
			{MenuID: 1},
			{MenuID: 2},
			{MenuID: 3},
		},
	}

	err := orderService.assignChef(chefs, order)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	for i, item := range order.Items {
		expectedChefID := chefs[i%len(chefs)].ID
		if item.ChefID != expectedChefID {
			t.Errorf("Expected ChefID %d, but got %d", expectedChefID, item.ChefID)
		}
	}
}

func TestProcessOrder(t *testing.T) {
	t.Parallel()

	orderService := setup()

	orderItem := entity.OrderItem{MenuID: 1}
	resultCh := make(chan entity.OrderItem, 1)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		orderService.processOrder(orderItem, resultCh)
	}()

	wg.Wait()

	result := <-resultCh

	if result.CreatedAt.IsZero() || result.FinishedAt.IsZero() || result.FinishedAt.Before(result.CreatedAt) {
		t.Error("Unexpected result from processOrder")
	}
}

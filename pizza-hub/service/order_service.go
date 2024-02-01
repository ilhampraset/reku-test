package service

import (
	"errors"
	"reku-code-test/pizza-hub/entity"
	"sync"
	"time"
)

type orderService struct {
	orderRepository entity.OrderRepository
	menuRepository  entity.MenuRepository
	chefRepository  entity.ChefRepository
}

func NewOrderService(
	orderRepository entity.OrderRepository,
	menuRepository entity.MenuRepository,
	chefRepository entity.ChefRepository) *orderService {
	return &orderService{
		orderRepository: orderRepository,
		menuRepository:  menuRepository,
		chefRepository:  chefRepository,
	}
}

func (s *orderService) CreateOrder(order *entity.Order) error {
	chefs, _ := s.chefRepository.All()
	s.assignChef(chefs, order)

	var wg sync.WaitGroup
	resultCh := make(chan entity.OrderItem, len(order.Items))
	processingOrderCh := make(chan struct{}, len(chefs))
	for _, item := range order.Items {
		wg.Add(1)
		go func(item entity.OrderItem) {
			defer wg.Done()
			processingOrderCh <- struct{}{}
			defer func() { <-processingOrderCh }()
			s.processOrder(item, resultCh)
		}(item)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()
	count := 0
	for result := range resultCh {
		order.Items[count].CreatedAt = result.CreatedAt
		order.Items[count].FinishedAt = result.FinishedAt
		count++
	}
	err := s.orderRepository.Create(order)
	if err != nil {
		return err
	}
	return nil
}

func (s *orderService) assignChef(chefs []entity.Chef, order *entity.Order) error {
	chefCount := len(chefs)
	if chefCount == 0 {
		return errors.New("no chefs available")
	}
	for i, _ := range order.Items {
		order.Items[i].ChefID = chefs[i%chefCount].ID
	}
	return nil
}

func (s *orderService) processOrder(orderItem entity.OrderItem, resultCh chan entity.OrderItem) {
	menus := s.menuRepository.All()
	orderItem.CreatedAt = time.Now()
	time.Sleep(time.Duration(menus[orderItem.MenuID-1].CookingTime) * time.Second)
	orderItem.FinishedAt = time.Now()
	result := orderItem
	resultCh <- result

}

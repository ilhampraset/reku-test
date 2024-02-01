package repository

import (
	"errors"
	"fmt"
	"reku-code-test/pizza-hub/entity"
	"sync"
)

type MemoryChefRepository struct {
	chefs map[string]*entity.Chef
	mu    sync.Mutex
}

func NewMemoryChefRepository() *MemoryChefRepository {
	return &MemoryChefRepository{
		chefs: make(map[string]*entity.Chef),
	}
}

// Add adds a new chef to the in-memory repository.
func (r *MemoryChefRepository) Add(chef *entity.Chef) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	chefID := fmt.Sprintf("%d", chef.ID)
	if _, exists := r.chefs[chefID]; exists {
		return errors.New("chef with the same ID already exists")
	}
	r.chefs[chefID] = chef
	return nil
}

func (r *MemoryChefRepository) All() ([]entity.Chef, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	chefs := make([]entity.Chef, 0, len(r.chefs))
	for _, chef := range r.chefs {
		chefs = append(chefs, *chef)
	}
	return chefs, nil
}

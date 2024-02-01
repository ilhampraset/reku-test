package repository

import (
	"errors"
	"fmt"
	"reku-code-test/pizza-hub/entity"
	"sync"
)

type MemoryMenuRepository struct {
	menus map[string]*entity.Menu
	mu    sync.Mutex
}

func NewMemoryMenuRepository() *MemoryMenuRepository {
	return &MemoryMenuRepository{
		menus: make(map[string]*entity.Menu),
	}
}
func (r *MemoryMenuRepository) Add(menu *entity.Menu) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	menuID := fmt.Sprintf("%d", menu.ID)
	if _, exists := r.menus[menuID]; exists {
		return errors.New("menu with the same ID already exists")
	}
	r.menus[menuID] = menu
	return nil
}

func (r *MemoryMenuRepository) All() []entity.Menu {
	r.mu.Lock()
	defer r.mu.Unlock()

	menus := make([]entity.Menu, 0, len(r.menus))
	for _, menu := range r.menus {
		menus = append(menus, *menu)
	}
	return menus
}

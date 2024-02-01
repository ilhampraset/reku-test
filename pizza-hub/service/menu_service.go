package service

import "reku-code-test/pizza-hub/entity"

type menuService struct {
	menuRepository entity.MenuRepository
}

func NewMenuService(menuRepository entity.MenuRepository) *menuService {
	return &menuService{
		menuRepository: menuRepository,
	}
}

func (s *menuService) AddMenu(menu *entity.Menu) error {
	return s.menuRepository.Add(menu)
}

func (s *menuService) GetAllMenu() ([]entity.Menu, error) {
	return s.menuRepository.All(), nil
}

package service

import (
	"reku-code-test/pizza-hub/entity"
)

type chefService struct {
	chefRepository entity.ChefRepository
}

func NewChefService(chefRepository entity.ChefRepository) *chefService {
	return &chefService{
		chefRepository: chefRepository,
	}
}

func (srv *chefService) AddNewChef(chef *entity.Chef) error {
	return srv.chefRepository.Add(chef)
}

func (srv *chefService) GetAllChef() ([]entity.Chef, error) {
	return srv.chefRepository.All()
}

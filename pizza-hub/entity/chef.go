package entity

type Chef struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type ChefRepository interface {
	Add(chef *Chef) error
	All() ([]Chef, error)
}

type ChefService interface {
	AddNewChef(chef *Chef) error
	GetAllChef() ([]Chef, error)
}

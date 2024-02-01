package entity

type Menu struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	CookingTime uint16 `json:"cooking_time"`
}

type MenuRepository interface {
	Add(menu *Menu) error
	All() []Menu
}

type MenuService interface {
	AddMenu(menu *Menu) error
	GetAllMenu() ([]Menu, error)
}

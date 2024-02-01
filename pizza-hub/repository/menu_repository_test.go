package repository

import (
	"reflect"
	"reku-code-test/pizza-hub/entity"
	"testing"
)

func TestMemoryMenuRepository_Add(t *testing.T) {
	repo := NewMemoryMenuRepository()

	menu := &entity.Menu{
		ID:          1,
		Name:        "Pizza Margherita",
		CookingTime: 3,
	}

	err := repo.Add(menu)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	err = repo.Add(menu)
	if err == nil {
		t.Error("Expected an error, but got none")
	}
}

func TestMemoryMenuRepository_All(t *testing.T) {
	repo := NewMemoryMenuRepository()

	menu1 := &entity.Menu{ID: 1, Name: "Pizza Margherita", CookingTime: 3}
	menu2 := &entity.Menu{ID: 2, Name: "Pepperoni Pizza", CookingTime: 5}
	repo.Add(menu1)
	repo.Add(menu2)

	expectedMenus := []entity.Menu{*menu1, *menu2}

	menus := repo.All()

	if !reflect.DeepEqual(menus, expectedMenus) {
		t.Errorf("Expected menus: %v, but got: %v", expectedMenus, menus)
	}
}

package repository

import (
	"reflect"
	"reku-code-test/pizza-hub/entity"
	"testing"
)

func TestMemoryChefRepository_Add(t *testing.T) {
	repo := NewMemoryChefRepository()

	chef := &entity.Chef{
		ID:   uint64(1),
		Name: "John Doe",
	}

	err := repo.Add(chef)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	err = repo.Add(chef)
	if err == nil {
		t.Error("Expected an error, but got none")
	}
}

func TestMemoryChefRepository_All(t *testing.T) {
	repo := NewMemoryChefRepository()

	chef1 := &entity.Chef{ID: 1, Name: "John Doe"}
	chef2 := &entity.Chef{ID: 2, Name: "Jane Doe"}
	repo.Add(chef1)
	repo.Add(chef2)

	expectedChefs := []entity.Chef{*chef1, *chef2}

	chefs, err := repo.All()

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if !reflect.DeepEqual(chefs, expectedChefs) {
		t.Errorf("Expected chefs: %v, but got: %v", expectedChefs, chefs)
	}
}

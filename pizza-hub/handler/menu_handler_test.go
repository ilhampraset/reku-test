package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reku-code-test/pizza-hub/entity"
	"testing"
)

type MockMenuService struct{}

func (m *MockMenuService) AddMenu(menu *entity.Menu) error {
	return nil
}

func (m *MockMenuService) GetAllMenu() ([]entity.Menu, error) {
	menus := []entity.Menu{
		{ID: 1, Name: "Pizza Cheese", CookingTime: 2},
		{ID: 2, Name: "Pizza Meatball", CookingTime: 3},
	}
	return menus, nil
}

func TestMenuHandler_AddMenu(t *testing.T) {
	// Create a new instance of the MenuHandler with the mock service
	handler := NewMenuHandler(&MockMenuService{})

	// Create a mock HTTP request payload
	newMenu := entity.Menu{ID: 1, Name: "Test Menu", CookingTime: 12}
	payload, _ := json.Marshal(newMenu)

	req, err := http.NewRequest(http.MethodPost, "/menus/add", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()

	handler.AddMenu(rec, req)

	// Check the status code
	if status := rec.Code; status != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, status)
	}
}

func TestMenuHandler_GetAllMenu(t *testing.T) {

	handler := NewMenuHandler(&MockMenuService{})
	req, err := http.NewRequest(http.MethodGet, "/menus", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler.GetAllMenu(rec, req)

	if status := rec.Code; status != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, status)
	}
}

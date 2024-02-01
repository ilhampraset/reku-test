package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reku-code-test/pizza-hub/entity"
	"testing"
)

type MockChefService struct{}

func (m *MockChefService) AddNewChef(chef *entity.Chef) error {
	return nil
}

func (m *MockChefService) GetAllChef() ([]entity.Chef, error) {
	chefs := []entity.Chef{
		{ID: 1, Name: "Chef 1"},
		{ID: 2, Name: "Chef 2"},
	}
	return chefs, nil
}

func TestChefHandler_AddChef(t *testing.T) {

	handler := NewChefHandler(&MockChefService{})

	newChef := entity.Chef{ID: 1, Name: "Test Chef"}
	payload, _ := json.Marshal(newChef)

	req, err := http.NewRequest(http.MethodPost, "/chefs/add", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()

	handler.AddChef(rec, req)

	if status := rec.Code; status != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, status)
	}
}

func TestChefHandler_GetAllChef(t *testing.T) {

	handler := NewChefHandler(&MockChefService{})

	req, err := http.NewRequest(http.MethodGet, "/chefs", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()

	handler.GetAllChef(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, status)
	}
}

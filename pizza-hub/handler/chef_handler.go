package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reku-code-test/pizza-hub/entity"
	"reku-code-test/pizza-hub/utils"
)

type ChefHandler struct {
	service entity.ChefService
}

func NewChefHandler(service entity.ChefService) *ChefHandler {
	return &ChefHandler{
		service: service,
	}
}

// AddChef handles the HTTP request to add a new chef.
func (h *ChefHandler) AddChef(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newChef entity.Chef
	if err := json.NewDecoder(r.Body).Decode(&newChef); err != nil {
		http.Error(w, "Invalid chef format", http.StatusBadRequest)
		return
	}

	if err := h.service.AddNewChef(&newChef); err != nil {
		http.Error(w, fmt.Sprintf("Failed to add chef: %v", err), http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, newChef, http.StatusCreated)
}

func (h *ChefHandler) GetAllChef(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	chefs, err := h.service.GetAllChef()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get chefs: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	utils.JsonResponse(w, chefs, http.StatusOK)
}

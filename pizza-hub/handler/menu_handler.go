package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reku-code-test/pizza-hub/entity"
	"reku-code-test/pizza-hub/utils"
)

type MenuHandler struct {
	service entity.MenuService
}

// NewMenuHandler creates a new instance of MenuHandler.
func NewMenuHandler(service entity.MenuService) *MenuHandler {
	return &MenuHandler{
		service: service,
	}
}

func (h *MenuHandler) AddMenu(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newMenu entity.Menu
	if err := json.NewDecoder(r.Body).Decode(&newMenu); err != nil {
		http.Error(w, "Invalid menu format", http.StatusBadRequest)
		return
	}

	if err := h.service.AddMenu(&newMenu); err != nil {
		http.Error(w, fmt.Sprintf("Failed to add menu: %v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	utils.JsonResponse(w, newMenu, http.StatusCreated)
}

func (h *MenuHandler) GetAllMenu(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	menus, err := h.service.GetAllMenu()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get menus: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	utils.JsonResponse(w, menus, http.StatusCreated)
}

package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reku-code-test/pizza-hub/entity"
	"reku-code-test/pizza-hub/utils"
	"time"
)

type OrderHandler struct {
	serviceOrder entity.OrderService
	serviceChef  entity.ChefService
	serviceMenu  entity.MenuService
}
type OrderRequest struct {
	ID    uint64 `json:"id"`
	Items []Item `json:"items"`
}

type OrderResponse struct {
	ID    uint64         `json:"id"`
	Items []ItemResponse `json:"items"`
}
type Item struct {
	MenuID uint64 `json:"menu_id"`
}

type ItemResponse struct {
	Menu       string    `json:"menu"`
	Chef       string    `json:"chef"`
	CreatedAt  time.Time `json:"created_at"`
	FinishedAt time.Time `json:"finished_at"`
}

func NewOrderHandler(serviceOrder entity.OrderService,
	serviceChef entity.ChefService,
	serviceMenu entity.MenuService) *OrderHandler {
	return &OrderHandler{
		serviceOrder: serviceOrder,
		serviceChef:  serviceChef,
		serviceMenu:  serviceMenu,
	}
}

func (h *OrderHandler) AddOrder(w http.ResponseWriter, r *http.Request) {
	var orderReq OrderRequest
	var orderResp OrderResponse
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&orderReq); err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid order format", http.StatusBadRequest)
		return
	}

	if len(orderReq.Items) == 0 {
		http.Error(w, "At least one item is required", http.StatusBadRequest)
		return
	}

	var newOrder entity.Order
	newOrder.ID = orderReq.ID
	for _, item := range orderReq.Items {
		newOrder.Items = append(newOrder.Items, entity.OrderItem{MenuID: item.MenuID})
	}

	if err := h.serviceOrder.CreateOrder(&newOrder); err != nil {
		http.Error(w, fmt.Sprintf("Failed to add order: %v", err), http.StatusInternalServerError)
		return
	}
	chefs, _ := h.serviceChef.GetAllChef()
	if len(chefs) == 0 {
		http.Error(w, "no chefs available to process orders", http.StatusBadRequest)
		return
	}
	chefMap := make(map[uint64]string)
	for _, chef := range chefs {
		chefMap[chef.ID] = chef.Name
	}
	menus, _ := h.serviceMenu.GetAllMenu()
	if len(menus) == 0 {
		http.Error(w, "No menus available", http.StatusBadRequest)
		return
	}
	menuMap := make(map[uint64]string)
	for _, menu := range menus {
		menuMap[menu.ID] = menu.Name
	}
	orderResp.ID = newOrder.ID
	for _, item := range newOrder.Items {
		chefName, chefExists := chefMap[item.ChefID]
		if !chefExists {
			http.Error(w, fmt.Sprintf("Chef with ID %d not found", item.ChefID), http.StatusInternalServerError)
			return
		}
		menuName, menuExists := menuMap[item.MenuID]
		if !menuExists {
			http.Error(w, "Menu with ID %d not found", int(item.MenuID))
			return
		}

		orderResp.Items = append(orderResp.Items, ItemResponse{
			Menu:       menuName,
			Chef:       chefName,
			CreatedAt:  item.CreatedAt,
			FinishedAt: item.FinishedAt,
		})
	}

	utils.JsonResponse(w, orderResp, http.StatusCreated)
}

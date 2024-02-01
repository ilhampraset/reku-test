// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"reku-code-test/pizza-hub/handler"
	"reku-code-test/pizza-hub/repository"
	"reku-code-test/pizza-hub/service"
)

func main() {

	chefRepo := repository.NewMemoryChefRepository()
	chefService := service.NewChefService(chefRepo)
	menuRepo := repository.NewMemoryMenuRepository()
	menuService := service.NewMenuService(menuRepo)
	orderRepo := repository.NewMemoryOrderRepository()
	orderService := service.NewOrderService(orderRepo, menuRepo, chefRepo)

	http.HandleFunc("/orders", handler.NewOrderHandler(orderService, chefService, menuService).AddOrder)
	http.HandleFunc("/chefs/add", handler.NewChefHandler(chefService).AddChef)
	http.HandleFunc("/chefs", handler.NewChefHandler(chefService).GetAllChef)
	http.HandleFunc("/menus", handler.NewMenuHandler(menuService).GetAllMenu)
	http.HandleFunc("/menus/add", handler.NewMenuHandler(menuService).AddMenu)

	fmt.Println("PizzaHub is serving at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

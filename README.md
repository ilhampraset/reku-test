# URL Shortener
A simple URL shortener service that allows users to create short aliases for long URLs.

## Installation

```bash
go run url-shortener/main.go
```

## Endpoint
Endpoint: GET /short-urls

Endpoint: POST /short-url
```
{
	"target_url" : "https://www.youtube.com",
	"expiry_date" : "2025-01-02 23:59:59"
}
```


# Pizza Hub API
The Pizza Hub API provides functionality to manage orders, chefs, and menus in a pizza restaurant.

## Installation

```bash
go run pizza-hub/main.go
```

## Resources
### 1. Chef
Endpoint: GET /chefs

Endpoint: POST /chefs/add
```
{
	"id" : 2,
	"name" : "john 2"
}
```

### 2. Menu 
Endpoint: GET /menus

Endpoint: POST /menus/add
```
{
	"id" : 2,
	"name" : "Pizza Chees",
	"cooking_time" : 3
}
```
### 1. Order
   Endpoint: POST /orders
```
{
  "id": 1,
  "items": [
    {
      "menu_id": 1,
    },
    {
      "menu_id": 2
    }
  ]
}
```

## Additional

```bash
go mod download
```

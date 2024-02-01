// cmd/main.go
package main

import (
	"fmt"
	"net/http"
	"reku-code-test/url-shortener/handler"
	"reku-code-test/url-shortener/repository"
	"reku-code-test/url-shortener/service"
)

func main() {
	urlStore := repository.NewMemoryStore()
	urlService := service.NewURLService(urlStore)
	urlHandler := handler.NewURLHandler(urlService)

	http.HandleFunc("/short-url", urlHandler.CreateShortURL)
	http.HandleFunc("/", urlHandler.RedirectToTargetURL)
	http.HandleFunc("/short-urls", urlHandler.GetAllURLs)

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}

package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reku-code-test/url-shortener/config"
	"reku-code-test/url-shortener/entity"
	"reku-code-test/url-shortener/service"
	"strconv"
	"strings"
	"time"
)

type UrlRequest struct {
	TargetURL  string `json:"target_url"`
	ExpiryDate string `json:"expiry_date"`
}

type UrlResponse struct {
	ShortURL   string    `json:"short_url"`
	TargetURL  string    `json:"target_url"`
	ExpiryDate time.Time `json:"expiry_date"`
	Counter    uint64    `json:"counter"`
}
type URLHandler struct {
	service *service.UrlService
}

func NewURLHandler(service *service.UrlService) *URLHandler {
	return &URLHandler{service: service}
}

func (h *URLHandler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	var reqData UrlRequest

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	layout := "2006-01-02 15:04:05"

	parsedTime, err := time.Parse(layout, reqData.ExpiryDate)

	if err != nil {
		return
	}
	shortURL, err := h.service.CreateShortURL(&entity.Url{TargetURL: reqData.TargetURL, ExpiryDate: parsedTime})
	if err != nil {
		http.Error(w, "Invalid target URL", http.StatusBadRequest)
		return
	}

	respData := UrlResponse{
		ShortURL:   fmt.Sprintf("%s/%s", config.HOST, shortURL),
		TargetURL:  reqData.TargetURL,
		ExpiryDate: parsedTime,
	}
	respondJSON(w, respData, http.StatusOK)
}

func (h *URLHandler) RedirectToTargetURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	shortURL := strings.TrimPrefix(r.URL.Path, "/")

	urlData, err := h.service.GeByShortUrl(shortURL)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	currentTime := time.Now()
	if currentTime.After(urlData.ExpiryDate) {
		http.Error(w, "Short URL has expired", http.StatusGone)
		return
	}

	urlData.ClickCount++
	http.Redirect(w, r, urlData.TargetURL, http.StatusSeeOther)
}

func (h *URLHandler) GetAllURLs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	descendingParam := r.URL.Query().Get("desc")
	descending, err := strconv.ParseBool(descendingParam)
	if err != nil {
		// Handle the error, e.g., default to false or return an error response
		descending = false
	}

	urls, err := h.service.GetAllURLs(descending)
	if err != nil {
		http.Error(w, "Invalid target URL", http.StatusBadRequest)
		return
	}
	respondJSON(w, urls, http.StatusOK)
}

func respondJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {

		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
	}
}

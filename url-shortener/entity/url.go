package entity

import "time"

type Url struct {
	ShortURL   string    `json:"short_url"`
	TargetURL  string    `json:"target_url"`
	ClickCount int       `json:"click_count"`
	ExpiryDate time.Time `json:"expiry_date"`
}

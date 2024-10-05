package domain

import "time"

type Session struct {
	ID         string    `json:"id"`
	CustomerID string    `json:"customer_id"`
	ExpiresAt  time.Time `json:"expires_at"`
}

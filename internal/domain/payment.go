package domain

import "time"

type Payment struct {
	ID         string    `json:"id"`
	CustomerID string    `json:"customer_id"`
	MerchantID string    `json:"merchant_id"`
	Amount     float64   `json:"amount"`
	Timestamp  time.Time `json:"timestamp"`
}

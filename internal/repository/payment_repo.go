package repository

import (
	"encoding/json"
	"os"
	"sync"

	"simple_bank/internal/domain"
)

type PaymentRepository struct {
	payments []domain.Payment
	mu       sync.RWMutex
}

func NewPaymentRepository() (*PaymentRepository, error) {
	repo := &PaymentRepository{}

	file, err := os.ReadFile("data/payments.json")
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(file, &repo.payments); err != nil {
		return nil, err
	}

	return repo, nil
}

func (r *PaymentRepository) Add(payment domain.Payment) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.payments = append(r.payments, payment)

	// Save to file
	file, err := json.MarshalIndent(r.payments, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile("data/payments.json", file, 0644)
}

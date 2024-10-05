package repository

import (
	"encoding/json"
	"os"
	"sync"

	"simple_bank/internal/domain"
)

type MerchantRepository struct {
	merchants map[string]domain.Merchant
	mu        sync.RWMutex
}

func NewMerchantRepository() (*MerchantRepository, error) {
	repo := &MerchantRepository{
		merchants: make(map[string]domain.Merchant),
	}

	file, err := os.ReadFile("data/merchants.json")
	if err != nil {
		return nil, err
	}

	var merchants []domain.Merchant
	if err := json.Unmarshal(file, &merchants); err != nil {
		return nil, err
	}

	for _, merchant := range merchants {
		repo.merchants[merchant.ID] = merchant
	}

	return repo, nil
}

func (r *MerchantRepository) GetByID(id string) (*domain.Merchant, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	merchant, exists := r.merchants[id]
	if !exists {
		return nil, false
	}
	return &merchant, true
}

func (r *MerchantRepository) Update(merchant domain.Merchant) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.merchants[merchant.ID] = merchant

	merchantsSlice := make([]domain.Merchant, 0, len(r.merchants))
	for _, m := range r.merchants {
		merchantsSlice = append(merchantsSlice, m)
	}

	file, err := json.MarshalIndent(merchantsSlice, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile("data/merchants.json", file, 0644)
}

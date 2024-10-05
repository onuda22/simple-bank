package repository

import (
	"encoding/json"
	"os"
	"sync"

	"simple_bank/internal/domain"
)

type CustomerRepository struct {
	customers map[string]domain.Customer
	mu        sync.RWMutex
}

func NewCustomerRepository() (*CustomerRepository, error) {
	repo := &CustomerRepository{
		customers: make(map[string]domain.Customer),
	}

	file, err := os.ReadFile("data/customers.json")
	if err != nil {
		return nil, err
	}

	var customers []domain.Customer
	if err := json.Unmarshal(file, &customers); err != nil {
		return nil, err
	}

	for _, customer := range customers {
		repo.customers[customer.ID] = customer
	}

	return repo, nil
}

func (r *CustomerRepository) GetByUsername(username string) (*domain.Customer, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, customer := range r.customers {
		if customer.Username == username {
			return &customer, true
		}
	}

	return nil, false
}

func (r *CustomerRepository) GetByID(id string) (*domain.Customer, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	customer, exists := r.customers[id]
	if !exists {
		return nil, false
	}
	return &customer, true
}

func (r *CustomerRepository) Update(customer domain.Customer) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.customers[customer.ID] = customer

	customersSlice := make([]domain.Customer, 0, len(r.customers))
	for _, c := range r.customers {
		customersSlice = append(customersSlice, c)
	}

	file, err := json.MarshalIndent(customersSlice, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile("data/customers.json", file, 0644)
}

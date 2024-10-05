package repository

import (
	"encoding/json"
	"os"
	"sync"

	"simple_bank/internal/domain"
)

type HistoryRepository struct {
	history []domain.History
	mu      sync.RWMutex
}

func NewHistoryRepository() (*HistoryRepository, error) {
	repo := &HistoryRepository{}

	file, err := os.ReadFile("data/history.json")
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(file, &repo.history); err != nil {
		return nil, err
	}

	return repo, nil
}

func (r *HistoryRepository) Add(entry domain.History) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.history = append(r.history, entry)

	file, err := json.MarshalIndent(r.history, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile("data/history.json", file, 0644)
}

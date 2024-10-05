package usecase

import (
	"errors"
	"time"

	"simple_bank/internal/domain"
	"simple_bank/internal/repository"
)

type LoginUseCase struct {
	customerRepo *repository.CustomerRepository
	historyRepo  *repository.HistoryRepository
}

func NewLoginUseCase(customerRepo *repository.CustomerRepository, historyRepo *repository.HistoryRepository) *LoginUseCase {
	return &LoginUseCase{
		customerRepo: customerRepo,
		historyRepo:  historyRepo,
	}
}

func (u *LoginUseCase) Login(username, password string) (*domain.Customer, error) {
	customer, exists := u.customerRepo.GetByUsername(username)
	if !exists || customer.Password != password {
		return nil, errors.New("invalid credentials")
	}

	// Log login activity
	u.historyRepo.Add(domain.History{
		ID:        generateID(),
		UserID:    customer.ID,
		Action:    "login",
		Details:   "User logged in",
		Timestamp: time.Now(),
	})

	return customer, nil
}

func generateID() string {
	return time.Now().Format("20060102150405")
}

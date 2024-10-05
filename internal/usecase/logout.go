package usecase

import (
	"time"

	"simple_bank/internal/domain"
	"simple_bank/internal/repository"
)

type LogoutUseCase struct {
	historyRepo *repository.HistoryRepository
}

func NewLogoutUseCase(historyRepo *repository.HistoryRepository) *LogoutUseCase {
	return &LogoutUseCase{
		historyRepo: historyRepo,
	}
}

func (u *LogoutUseCase) Logout(customerID string) error {

	u.historyRepo.Add(domain.History{
		ID:        generateID(),
		UserID:    customerID,
		Action:    "logout",
		Details:   "User logged out",
		Timestamp: time.Now(),
	})

	return nil
}

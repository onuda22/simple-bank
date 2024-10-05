package usecase

import (
	"errors"
	"time"

	"simple_bank/internal/domain"
	"simple_bank/internal/repository"

	"github.com/dgrijalva/jwt-go"
)

type LoginUseCase struct {
	customerRepo *repository.CustomerRepository
	historyRepo  *repository.HistoryRepository
	jwtSecret    []byte
}

func NewLoginUseCase(customerRepo *repository.CustomerRepository, historyRepo *repository.HistoryRepository, jwtSecret string) *LoginUseCase {
	return &LoginUseCase{
		customerRepo: customerRepo,
		historyRepo:  historyRepo,
		jwtSecret:    []byte(jwtSecret),
	}
}

func (u *LoginUseCase) Login(username, password string) (string, error) {
	customer, exists := u.customerRepo.GetByUsername(username)
	if !exists || customer.Password != password {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"customer_id": customer.ID,
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(u.jwtSecret)
	if err != nil {
		return "", err
	}

	// Log login activity
	u.historyRepo.Add(domain.History{
		ID:        generateID(),
		UserID:    customer.ID,
		Action:    "login",
		Details:   "User logged in",
		Timestamp: time.Now(),
	})

	return tokenString, nil
}

func generateID() string {
	return time.Now().Format("20060102150405")
}

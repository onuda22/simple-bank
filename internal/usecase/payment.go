package usecase

import (
	"errors"
	"time"

	"simple_bank/internal/domain"
	"simple_bank/internal/repository"
)

type PaymentUseCase struct {
	customerRepo *repository.CustomerRepository
	merchantRepo *repository.MerchantRepository
	paymentRepo  *repository.PaymentRepository
	historyRepo  *repository.HistoryRepository
}

func NewPaymentUseCase(customerRepo *repository.CustomerRepository, merchantRepo *repository.MerchantRepository, paymentRepo *repository.PaymentRepository, historyRepo *repository.HistoryRepository) *PaymentUseCase {
	return &PaymentUseCase{
		customerRepo: customerRepo,
		merchantRepo: merchantRepo,
		paymentRepo:  paymentRepo,
		historyRepo:  historyRepo,
	}
}

func (u *PaymentUseCase) MakePayment(customerID, merchantID string, amount float64) error {
	customer, exists := u.customerRepo.GetByID(customerID)
	if !exists {
		return errors.New("customer not found")
	}

	merchant, exists := u.merchantRepo.GetByID(merchantID)
	if !exists {
		return errors.New("merchant not found")
	}

	if customer.Balance < amount {
		return errors.New("insufficient funds")
	}

	// Update balances
	customer.Balance -= amount
	merchant.Balance += amount

	// Save changes
	if err := u.customerRepo.Update(*customer); err != nil {
		return err
	}
	if err := u.merchantRepo.Update(*merchant); err != nil {
		return err
	}

	// Record payment
	payment := domain.Payment{
		ID:         generateID(),
		CustomerID: customerID,
		MerchantID: merchantID,
		Amount:     amount,
		Timestamp:  time.Now(),
	}
	if err := u.paymentRepo.Add(payment); err != nil {
		return err
	}

	// Log payment activity
	u.historyRepo.Add(domain.History{
		ID:        generateID(),
		UserID:    customerID,
		Action:    "payment",
		Details:   "Payment made to merchant",
		Timestamp: time.Now(),
	})

	return nil
}

package main

import (
	"log"
	"net/http"

	"simple_bank/internal/api"
	"simple_bank/internal/api/handlers"
	"simple_bank/internal/repository"
	"simple_bank/internal/usecase"
)

func main() {
	// Initialize repositories
	customerRepo, err := repository.NewCustomerRepository()
	if err != nil {
		log.Fatalf("Failed to initialize customer repository: %v", err)
	}

	merchantRepo, err := repository.NewMerchantRepository()
	if err != nil {
		log.Fatalf("Failed to initialize merchant repository: %v", err)
	}

	paymentRepo, err := repository.NewPaymentRepository()
	if err != nil {
		log.Fatalf("Failed to initialize payment repository: %v", err)
	}

	historyRepo, err := repository.NewHistoryRepository()
	if err != nil {
		log.Fatalf("Failed to initialize history repository: %v", err)
	}

	// Initialize use cases
	loginUseCase := usecase.NewLoginUseCase(customerRepo, historyRepo)
	paymentUseCase := usecase.NewPaymentUseCase(customerRepo, merchantRepo, paymentRepo, historyRepo)
	logoutUseCase := usecase.NewLogoutUseCase(historyRepo)

	// Initialize handlers
	loginHandler := handlers.NewLoginHandler(loginUseCase)
	paymentHandler := handlers.NewPaymentHandler(paymentUseCase)
	logoutHandler := handlers.NewLogoutHandler(logoutUseCase)

	// Setup routes
	router := api.SetupRoutes(loginHandler, paymentHandler, logoutHandler)

	// Start server
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

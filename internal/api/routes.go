package api

import (
	"net/http"

	"simple_bank/internal/api/handlers"
	"simple_bank/internal/middleware"
)

func SetupRoutes(
	loginHandler *handlers.LoginHandler,
	paymentHandler *handlers.PaymentHandler,
	logoutHandler *handlers.LogoutHandler,
) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/login", loginHandler.Login)
	mux.HandleFunc("/payment", middleware.Authenticate(paymentHandler.MakePayment))
	mux.HandleFunc("/logout", middleware.Authenticate(logoutHandler.Logout))

	return mux
}

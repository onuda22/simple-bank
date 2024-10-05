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
	jwtSecret []byte,
) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/login", loginHandler.Login)
	mux.HandleFunc("/payment", middleware.Authenticate(jwtSecret)(paymentHandler.MakePayment))
	mux.HandleFunc("/logout", middleware.Authenticate(jwtSecret)(logoutHandler.Logout))

	return mux
}

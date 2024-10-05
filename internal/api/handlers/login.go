package handlers

import (
	"encoding/json"
	"net/http"

	"simple_bank/internal/usecase"
)

type LoginHandler struct {
	loginUseCase *usecase.LoginUseCase
}

func NewLoginHandler(loginUseCase *usecase.LoginUseCase) *LoginHandler {
	return &LoginHandler{loginUseCase: loginUseCase}
}

func (h *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	customer, err := h.loginUseCase.Login(req.Username, req.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Login successful", "customer_id": customer.ID})
}

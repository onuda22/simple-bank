package handlers

import (
	"encoding/json"
	"net/http"

	"simple_bank/internal/usecase"
)

type LogoutHandler struct {
	logoutUseCase *usecase.LogoutUseCase
}

func NewLogoutHandler(logoutUseCase *usecase.LogoutUseCase) *LogoutHandler {
	return &LogoutHandler{logoutUseCase: logoutUseCase}
}

func (h *LogoutHandler) Logout(w http.ResponseWriter, r *http.Request) {
	customerID := r.Header.Get("CustomerID")
	if customerID == "" {
		http.Error(w, "Customer ID not found", http.StatusBadRequest)
		return
	}

	err := h.logoutUseCase.Logout(customerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Logout successful"})
}

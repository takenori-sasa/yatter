package statuses

import (
	"encoding/json"
	"fmt"
	"net/http"

	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/auth"
)

// Request body for `POST /v1/accounts`
type PostStatusInput struct {
	Content string
}

// Handle request for `POST /v1/accounts`
func (h *handler) CreateStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req PostStatusInput
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}
	account := auth.AccountOf(r)
	if account == nil {
		http.Error(w, "Authorization required", http.StatusUnauthorized)
		return
	}

	status := new(object.Status)
	status.Content = &req.Content
	status.AccountID = account.ID

	res, err := h.sr.CreateStatus(ctx, status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

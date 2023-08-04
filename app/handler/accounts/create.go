package accounts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"yatter-backend-go/app/domain/object"
)

// Request body for `POST /v1/accounts`
type AddRequest struct {
	Username string
	Password string
}

// Handle request for `POST /v1/accounts`
// Create function is responsible for creating a new account.
// It expects a username and password in the request body.
// If the account is created successfully, it returns a JSON representation of the account.
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Parse request body
	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// Invalid request body
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate request body. Username and password are required.
	if req.Username == "" || req.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	// Initialize new account object and set username and password
	account := new(object.Account)
	account.Username = req.Username
	if err := account.SetPassword(req.Password); err != nil {
		// Failed to set password
		http.Error(w, fmt.Sprintf("Failed to set password: %v", err.Error()), http.StatusInternalServerError)
		return
	}

	// Create new user account
	res, err := h.ar.CreateUser(ctx, account)
	if err != nil {
		// Failed to create user
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set response header
	w.Header().Set("Content-Type", "application/json")

	// Encode and send response
	if err := json.NewEncoder(w).Encode(res); err != nil {
		// Failed to encode the response
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

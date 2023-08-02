package statuses

import (
	"encoding/json"
	"net/http"

	"yatter-backend-go/app/domain/object"
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	status := new(object.Status)
	status.Content = &req.Content
	// if err := account.SetPassword(req.Password); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// panic("Must Implement Account Registration")
	res, err := h.ar.CreateStatus(ctx, status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

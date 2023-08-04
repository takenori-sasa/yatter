package statuses

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// New DeleteStatus handler
func (h *handler) DeleteStatus(w http.ResponseWriter, r *http.Request) {
	statusID := chi.URLParam(r, "id")

	// You may need to convert statusID to the type expected by DeleteStatus (e.g., int or int64)
	// Convert string to int64
	statusIDInt64, err := strconv.ParseInt(statusID, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid status ID: %v", err), http.StatusBadRequest)
		return
	}

	if err := h.sr.DeleteStatus(r.Context(), statusIDInt64); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

package timelines

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Handle request for `Get /v1/statuses`
func (h *handler) GetHomeTimeline(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	IDParam := chi.URLParam(r, "id")
	statusID, err := strconv.ParseInt(IDParam, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// int64の検査
	res, err := h.ar.FindStatus(ctx, statusID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

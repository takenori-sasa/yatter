package timelines

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Handle request for `Get /v1/statuses`
func (h *handler) GetPublicTimeline(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	MaxIDParam := chi.URLParam(r, "max_id")
	MaxID, err := strconv.ParseInt(MaxIDParam, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	SinceIDParam := chi.URLParam(r, "since_id")
	SinceID, err := strconv.ParseInt(SinceIDParam, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	LimitParam := chi.URLParam(r, "limit")
	Limit, err := strconv.ParseInt(LimitParam, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// int64の検査
	res, err := h.tr.FindPublicTimeline(ctx, MaxID, SinceID, Limit)

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

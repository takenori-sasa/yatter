package timelines

import (
	"encoding/json"
	"net/http"
	"strconv"
	"yatter-backend-go/app/handler/auth"
)

// Handle request for `Get /v1/statuses`
func (h *handler) GetHomeTimeline(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	MaxIDParam := r.URL.Query().Get("max_id")
	if MaxIDParam == "" {
		MaxIDParam = "-1"
	}
	MaxID, err := strconv.ParseInt(MaxIDParam, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	SinceIDParam := r.URL.Query().Get("since_id")
	if SinceIDParam == "" {
		SinceIDParam = "-1"
	}
	SinceID, err := strconv.ParseInt(SinceIDParam, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	LimitParam := r.URL.Query().Get("limit")
	if LimitParam == "" {
		LimitParam = "40"
	}
	Limit, err := strconv.ParseInt(LimitParam, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	account := auth.AccountOf(r)
	timeline, err := h.tr.FindHomeTimeline(ctx, MaxID, SinceID, Limit, account)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(timeline); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

package timelines

import (
	"encoding/json"
	"net/http"
	"strconv"
	"yatter-backend-go/app/handler/auth"
)

// Handle request for `Get /v1/statuses`
// This function extracts parameters (max_id, since_id, limit) from the request,
// validates and uses them to retrieve home timeline from the database.
// In case of error, it responds with the respective error message and status code.
func (h *handler) GetHomeTimeline(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Parse and validate max_id parameter
	MaxIDParam := r.URL.Query().Get("max_id")
	if MaxIDParam == "" {
		MaxIDParam = "-1"
	}
	MaxID, err := strconv.ParseInt(MaxIDParam, 10, 64)
	if err != nil {
		http.Error(w, "Invalid max_id parameter", http.StatusBadRequest)
		return
	}

	// Parse and validate since_id parameter
	SinceIDParam := r.URL.Query().Get("since_id")
	if SinceIDParam == "" {
		SinceIDParam = "-1"
	}
	SinceID, err := strconv.ParseInt(SinceIDParam, 10, 64)
	if err != nil {
		http.Error(w, "Invalid since_id parameter", http.StatusBadRequest)
		return
	}

	// Parse and validate limit parameter
	LimitParam := r.URL.Query().Get("limit")
	if LimitParam == "" {
		LimitParam = "40"
	}
	Limit, err := strconv.ParseInt(LimitParam, 10, 64)
	if err != nil {
		http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
		return
	}

	// Get account from the request context
	account := auth.AccountOf(r)

	// Call FindHomeTimeline with parsed parameters and handle possible error
	timeline, err := h.tr.FindHomeTimeline(ctx, MaxID, SinceID, Limit, account)
	if err != nil {
		http.Error(w, "Failed to retrieve home timeline", http.StatusInternalServerError)
		return
	}

	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode and write the timeline to the response body and handle possible error
	if err := json.NewEncoder(w).Encode(timeline); err != nil {
		http.Error(w, "Failed to encode timeline to JSON", http.StatusInternalServerError)
		return
	}
}

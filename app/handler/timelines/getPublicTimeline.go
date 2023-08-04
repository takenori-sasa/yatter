package timelines

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// Handle request for `Get /v1/statuses`
// This function extracts the max_id, since_id, and limit parameters from the request,
// converts them to the appropriate types, and then calls the FindPublicTimeline function.
// If an error occurs at any step, it sends a response with an appropriate HTTP status code and error message.
func (h *handler) GetPublicTimeline(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Extract and convert max_id parameter
	MaxIDParam := r.URL.Query().Get("max_id")
	if MaxIDParam == "" {
		MaxIDParam = "-1"
	}
	MaxID, err := strconv.ParseInt(MaxIDParam, 10, 64)
	if err != nil {
		http.Error(w, "Failed to parse max_id parameter", http.StatusBadRequest)
		return
	}

	// Extract and convert since_id parameter
	SinceIDParam := r.URL.Query().Get("since_id")
	if SinceIDParam == "" {
		SinceIDParam = "-1"
	}
	SinceID, err := strconv.ParseInt(SinceIDParam, 10, 64)
	if err != nil {
		http.Error(w, "Failed to parse since_id parameter", http.StatusBadRequest)
		return
	}

	// Extract and convert limit parameter
	LimitParam := r.URL.Query().Get("limit")
	if LimitParam == "" {
		LimitParam = "40"
	}
	Limit, err := strconv.ParseInt(LimitParam, 10, 64)
	if err != nil {
		http.Error(w, "Failed to parse limit parameter", http.StatusBadRequest)
		return
	}

	// Call FindPublicTimeline function
	timeline, err := h.tr.FindPublicTimeline(ctx, MaxID, SinceID, Limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode and send response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(timeline); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

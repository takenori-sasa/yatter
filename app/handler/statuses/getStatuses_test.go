package statuses_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"
	"yatter-backend-go/app/handler/statuses"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetStatusHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStatusRepo := repository.NewMockStatus(ctrl)

	handler := statuses.NewHandler(mockStatusRepo)

	statusID := int64(1) // Assuming this status ID exists
	req, _ := http.NewRequest("GET", "/v1/statuses/"+strconv.FormatInt(statusID, 10), nil)

	content := "test content"
	// set up the expected Status object and return values
	expectedStatus := &object.Status{
		ID:        statusID,
		Content:   &content,
		AccountID: 1, // assuming an account with ID 1 exists
	}
	mockStatusRepo.EXPECT().FindStatus(gomock.Any(), statusID).Return(expectedStatus, nil)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", strconv.FormatInt(statusID, 10))
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	// create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// call the handler
	handler.GetStatus(rr, req)

	// check the status code and body of the response
	assert.Equal(t, http.StatusOK, rr.Code)

	var responseStatus object.Status
	err := json.NewDecoder(rr.Body).Decode(&responseStatus)
	assert.NoError(t, err)
	assert.Equal(t, expectedStatus, &responseStatus)
}

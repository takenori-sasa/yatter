package statuses

import (
	"net/http"
	"yatter-backend-go/app/domain/repository"

	"github.com/go-chi/chi/v5"
)

// Implementation of handler
type handler struct {
	ar repository.Status
}

// Create Handler for `/v1/accounts/`
func NewRouter(ar repository.Status) http.Handler {
	r := chi.NewRouter()

	h := &handler{ar}
	r.Post("/", h.CreateStatus)
	r.Get("/{id}", h.GetStatuses)

	return r
}

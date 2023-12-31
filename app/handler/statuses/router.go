package statuses

import (
	"net/http"
	"yatter-backend-go/app/domain/repository"
	"yatter-backend-go/app/handler/auth"

	"github.com/go-chi/chi/v5"
)

// Implementation of handler
type handler struct {
	sr repository.Status
}

// NewHandler returns a new handler instance.
func NewHandler(sr repository.Status) *handler {
	return &handler{
		sr: sr,
	}
}

// Create Handler for `/v1/accounts/`
func NewRouter(sr repository.Status, ar repository.Account) http.Handler {
	r := chi.NewRouter()

	h := &handler{sr}
	r.Get("/{id}", h.GetStatus)
	r.Route("/", func(r chi.Router) {
		r.Use(auth.Middleware(ar))
		r.Post("/", h.CreateStatus)
		r.Delete("/{id}", h.DeleteStatus)
	})
	return r
}

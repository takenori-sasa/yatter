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

// Create Handler for `/v1/accounts/`
func NewRouter(sr repository.Status, ar repository.Account) http.Handler {
	r := chi.NewRouter()

	h := &handler{sr}
	r.Get("/{id}", h.GetStatus)
	r.Route("/", func(r chi.Router) {
		r.Use(auth.Middleware(ar))
		r.Post("/", h.CreateStatus)
	})
	return r
}

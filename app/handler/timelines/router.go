package timelines

import (
	"net/http"
	"yatter-backend-go/app/domain/repository"
	"yatter-backend-go/app/handler/auth"

	"github.com/go-chi/chi/v5"
)

// Implementation of handler
type handler struct {
	ar repository.Status
	tr repository.Timeline
}

// Create Handler for `/v1/accounts/`
func NewRouter(sr repository.Status, ar repository.Account, tr repository.Timeline) http.Handler {
	r := chi.NewRouter()

	h := &handler{sr, tr}
	r.Get("/public", h.GetPublicTimeline)
	r.Route("/", func(r chi.Router) {
		r.Use(auth.Middleware(ar))
		r.Get("/home", h.GetHomeTimeline)
	})
	return r
}

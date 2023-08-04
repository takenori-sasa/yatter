package timelines

import (
	"net/http"
	"yatter-backend-go/app/domain/repository"
	"yatter-backend-go/app/handler/auth"

	"github.com/go-chi/chi/v5"
)

// Implementation of handler
type handler struct {
	tr repository.Timeline
}

// NewRouter creates a new router for the `/v1/timelines/` route
func NewRouter(ar repository.Account, tr repository.Timeline) http.Handler {
	r := chi.NewRouter()

	h := &handler{tr}
	r.Get("/public", h.GetPublicTimeline)
	r.Route("/", func(r chi.Router) {
		r.Use(auth.Middleware(ar))
		r.Get("/home", h.GetHomeTimeline)
	})
	return r
}

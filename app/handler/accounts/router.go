package accounts

import (
	"net/http"
	"yatter-backend-go/app/domain/repository"
	"yatter-backend-go/app/handler/accounts/relationships"

	"github.com/go-chi/chi/v5"
)

// Implementation of handler
type handler struct {
	ar repository.Account
	rr repository.Relationship
}

// Create Handler for `/v1/accounts/`
func NewRouter(ar repository.Account, rr repository.Relationship) http.Handler {
	r := chi.NewRouter()

	h := &handler{ar, rr}
	r.Post("/", h.Create)
	r.Mount("/", relationships.NewRouter(rr))
	r.Get("/{username}", h.GetUserInfo)

	return r
}

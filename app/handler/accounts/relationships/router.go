package relationships

import (
	"net/http"
	"yatter-backend-go/app/domain/repository"

	"github.com/go-chi/chi/v5"
)

// Implementation of handler
type handler struct {
	rr repository.Relationship
}

// Create Handler for `/v1/accounts/`
func NewRouter(rr repository.Relationship) http.Handler {
	r := chi.NewRouter()

	// h := &handler{rr}

	return r
}

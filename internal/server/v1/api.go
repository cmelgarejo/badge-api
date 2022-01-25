package v1

import (
	"net/http"

	"github.com/cmelgarejo/badge-api/internal/data"
	"github.com/go-chi/chi"
)

// New returns the API V1 Handler with configuration.
func New() http.Handler {
	r := chi.NewRouter()

	ur := &UserRouter{
		Repository: &data.UserRepository{
			Data: data.New(),
		},
	}

	r.Mount("/users", ur.Routes())

	br := &BadgeRouter{
		Repository: &data.BadgeRepository{
			Data: data.New(),
		},
	}

	r.Mount("/badges", br.Routes())

	return r
}

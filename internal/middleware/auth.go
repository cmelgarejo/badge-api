package middleware

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/cmelgarejo/badge-api/pkg/claim"
	"github.com/cmelgarejo/badge-api/pkg/response"
)

type key string

// Context keys
const (
	UserIDKey key = "id"
)

// Authorize is a middleware that verifies if the token is valid.
func Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		tokenString, err := tokenFromAuthorization(authorization)
		if err != nil {
			_ = response.HTTPError(w, r, http.StatusUnauthorized, err.Error())
			return
		}

		c, err := claim.GetFromToken(tokenString, os.Getenv("JWT_SECRET"))
		if err != nil {
			_ = response.HTTPError(w, r, http.StatusUnauthorized, err.Error())
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, UserIDKey, c.ID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// tokenFromAuthorization extracts token from the authorization string.
func tokenFromAuthorization(authorization string) (string, error) {
	if authorization == "" {
		return "", errors.New("autorization is required")
	}

	if !strings.HasPrefix(authorization, "Bearer") {
		return "", errors.New("invalid autorization format")
	}

	l := strings.Split(authorization, " ")
	if len(l) != 2 {
		return "", errors.New("invalid autorization format")
	}

	return l[1], nil
}

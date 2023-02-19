package middlewares

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/otaxhu/notes-go-application/environment"
	"github.com/otaxhu/notes-go-application/models"
)

var (
	NO_AUTH_NEEDED = []string{
		"login",
		"signup",
	}
)

func shouldCheckAuth(route string) bool {
	for _, p := range NO_AUTH_NEEDED {
		if strings.Contains(route, p) {
			return false
		}
	}
	return true
}

func CheckAuth(f http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if !shouldCheckAuth(r.URL.Path) {
				f.ServeHTTP(w, r)
				return
			}
			tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
			if _, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(t *jwt.Token) (interface{}, error) {
				return []byte(environment.JWTSecret), nil
			}); err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			f.ServeHTTP(w, r)
		})
}

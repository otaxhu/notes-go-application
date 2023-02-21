package authorization

import (
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/otaxhu/notes-go-application/environment"
	"github.com/otaxhu/notes-go-application/models"
)

func GetClaims(tokenString string) (*models.AppClaims, error) {
	token, err := jwt.ParseWithClaims(strings.TrimSpace(tokenString), &models.AppClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(environment.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*models.AppClaims); !ok || !token.Valid {
		return nil, fmt.Errorf("unauthorized")
	} else {
		return claims, nil
	}
}

package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func isAuthorized(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "No token found", http.StatusUnauthorized)
			return
		}
		tokenString := strings.Split(authHeader, " ")[1]
		signingKey := []byte(os.Getenv("SECRET_KEY"))
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				http.Error(w, "Error parsing token", http.StatusInternalServerError)
				return nil, fmt.Errorf("Error parsing token")
			}
			return signingKey, nil
		})
		if err != nil {
			http.Error(w, "Token expired", http.StatusUnauthorized)
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			role := claims["role"]
			switch role {
			case "admin":
				r.Header.Set("role", "admin")
				return
			case "user":
				r.Header.Set("role", "user")
				return
			}
		}
		next.ServeHTTP(w, r)
	}
}

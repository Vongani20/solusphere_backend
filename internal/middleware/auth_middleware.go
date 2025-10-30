package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("SECRET_KEY_CHANGE_THIS")

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if !token.Valid {
			http.Error(w, "invalid/expired token", http.StatusUnauthorized)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		username := claims["username"].(string)

		ctx := context.WithValue(r.Context(), "user", username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

package middleware

import (
	"context"
	"encoding/json"
	"github.com/zerodays/sistem-auth/user"
	"net/http"
	"strings"
)

const contextUserKey = "sistem_auth_user"

// tokenFromRequest gets token from request. It first tries getting
// token from Authorization header. If there is no valid token present in Authorization header,
// function takes the token from GET parameters.
func tokenFromRequest(r *http.Request) string {
	// Get header value.
	header := r.Header.Get("Authorization")

	// Split header value to components.
	authComponents := strings.Split(header, " ")
	if len(authComponents) == 2 {
		// Extract auth type and token
		authType := authComponents[0]
		token := authComponents[1]

		// Check if auth type is bearer
		if authType == "Bearer" {
			return token
		}
	}

	// Get token from GET parameter.
	return r.URL.Query().Get("key")
}

// Middleware authenticates user from access token provided in Authorization header or GET parameter.
// User instance is added to request context for future handlers.
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get access token string from request.
		accessToken := tokenFromRequest(r)

		// Check if token is valid and get user from token.
		u, err := user.FromToken(accessToken)

		// Add user to context if valid.
		if err == nil {
			ctx := context.WithValue(r.Context(), contextUserKey, u)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

// UserFromRequest gets authenticated user from request context.
// If there is no user in request context, nil is returned.
func UserFromRequest(r *http.Request) *user.User {
	u, ok := r.Context().Value(contextUserKey).(*user.User)
	if ok {
		return u
	} else {
		return nil
	}
}

// RequiredMiddleware checks if user is authenticated before executing next handle.
// If user is not authenticated, error is written to response.
// This middleware has to be executed after auth.Middleware.
func RequiredMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u := UserFromRequest(r)
		if u == nil {
			// Write error to response.
			res, _ := json.Marshal(map[string]interface{}{
				"error_code": "invalid_token",
				"metadata":   nil,
			})

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write(res)
		} else {
			// Continue to next handler.
			next.ServeHTTP(w, r)
		}
	})
}

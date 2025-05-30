package middleware

import (
	"context"
	"ecom/shared"
	"net/http"

	"gorm.io/gorm"
)

var dbKey = shared.ContextKey("db")

// DatabaseMiddleware injects the database connection into the request context.
func DatabaseMiddleware(db *gorm.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, dbKey, db)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

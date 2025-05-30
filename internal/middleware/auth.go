package middleware

import (
	"ecom/internal/utils"
	"net/http"
	"os"
)

func AuthBearerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, exist := utils.ExtractBearerTokenFromHeader(r)
		if !exist || token == "" {
			utils.UnauthorizedHandler(w, r)
			return
		}

		// let's now validate the token
		valid, err := utils.ValidateJWTToken(token, os.Getenv("JWT_SECRET_KEY"))
		if !valid || err != nil {
			utils.UnauthorizedHandler(w, r)
			return
		}



		next.ServeHTTP(w, r) // Call the next handler if validation passes
	})
}

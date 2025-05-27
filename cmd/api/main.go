package main

import (
	"ecom/internal/utils"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Specific handler for GET /
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			utils.NotFoundHandler(w, r)
			return
		}
		utils.ReplyJson(w, http.StatusOK, map[string]string{"message": "Welcome to the E-commerce API"})
	})

	log.Panic(http.ListenAndServe(":3030", mux))
}

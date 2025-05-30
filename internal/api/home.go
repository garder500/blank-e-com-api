package api

import (
	"ecom/internal/utils"
	"net/http"
)

type HomeResponseFormat struct {
	Message string `json:"message"`
	Url     string `json:"url"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" || r.Method != http.MethodGet {
		utils.NotFoundHandler(w, r)
		return
	}

	utils.ReplyJson(w, http.StatusOK, HomeResponseFormat{
		Message: "Welcome to the E-commerce API",
		Url:     r.Host,
	})
}

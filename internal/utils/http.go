package utils

import (
	"encoding/json"
	"net/http"
)

type DetailsResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Stack   string `json:"stack,omitempty"`
}

type FormattedResponse struct {
	Message string            `json:"message"`
	Code    int               `json:"code"`
	Details []DetailsResponse `json:"details,omitempty"`
}

func ReplyJson(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func ReplyError(w http.ResponseWriter, status int, data FormattedResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	ReplyError(w, http.StatusNotFound, FormattedResponse{
		Message: "Resource not found",
		Code:    http.StatusNotFound,
		Details: []DetailsResponse{
			{
				Code:    http.StatusNotFound,
				Message: "The requested resource could not be found",
				Stack:   "Check the URL and try again",
			},
		},
	})
}

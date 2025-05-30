package api

import (
	"ecom/internal/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func HandleMux(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()
	// register the db as a context value if needed
	router.Use(middleware.DatabaseMiddleware(db))
	router.HandleFunc("/", HomeHandler)
	return router
}

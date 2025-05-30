package utils

import (
	"ecom/shared"
	"fmt"
	"net/http"

	gorm "gorm.io/gorm"
)

func GetDBFromReq(r *http.Request) (*gorm.DB, error) {
	db, ok := r.Context().Value(shared.ContextKey("db")).(*gorm.DB)
	if !ok || db == nil {
		return nil, fmt.Errorf("database not found in request context")
	}
	return db, nil
}

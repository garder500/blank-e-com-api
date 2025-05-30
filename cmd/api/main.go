package main

import (
	"ecom/internal/api"
	"ecom/internal/models"
	"ecom/internal/utils"
	"fmt"
	"log"
	"net/http"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {

	if err := utils.LoadEnv(utils.EnvVar{
		Key:          "PORT",
		DefaultValue: "3030",
	}, utils.EnvVar{
		Key:          "JWT_SECRET_KEY",
		DefaultValue: "your-secret-key",
	}); err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	// Set the log prefix
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.User{}, &models.Role{}, &models.Permissions{}, &models.ApiKey{})
	// Specific handler for GET /
	mux := api.HandleMux(db)

	log.Panic(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), mux))
}

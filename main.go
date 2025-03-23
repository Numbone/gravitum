package main

import (
	"context"
	"log"
	"net/http"

	"github.com/Numbone/user-api/config"
	"github.com/Numbone/user-api/internal/database"
	"github.com/Numbone/user-api/internal/handler"
	"github.com/Numbone/user-api/internal/repository"
	"github.com/Numbone/user-api/internal/service"
)

func main() {
	cfg := config.LoadConfig()

	ctx := context.Background()

	db, err := database.Connect(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	http.HandleFunc("/users", userHandler.CreateUser)
	http.HandleFunc("/users/", userHandler.GetUser)

	log.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

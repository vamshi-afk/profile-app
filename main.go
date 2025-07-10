package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"profile-app/database"
	"profile-app/handlers"
	"profile-app/middleware"
)

func main() {
	if err := database.Init(); err != nil {
		log.Fatalf("DB init failed: %v\n", err)
	}

	r := mux.NewRouter()

	// public routes
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

	// routes which can be accessed after login session creates
	protected := r.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthMiddleware)
	protected.HandleFunc("/profile", handlers.ProfileHandler)
	protected.HandleFunc("/logout", handlers.LogoutHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback for local dev
	}
	log.Printf("Server running on http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))

}

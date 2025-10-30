package main

import (
	"Solusphere/internal/db"
	"Solusphere/internal/handlers"
	"Solusphere/internal/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	db.Connect() // ← important

	r := mux.NewRouter()
	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.Handle("/profile", middleware.RequireAuth(http.HandlerFunc(handlers.Profile))).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5174"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	log.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

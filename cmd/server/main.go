package main

import (
	"log"
	"net/http"

	"Solusphere/internal/db"
	"Solusphere/internal/handlers"
	"Solusphere/internal/middleware"

	"github.com/gorilla/mux"
)

func main() {

	db.Connect("root", "", "127.0.0.1:3306", "solusphere")

	r := mux.NewRouter()

	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.Handle("/profile", middleware.RequireAuth(http.HandlerFunc(handlers.Profile))).Methods("GET")

	log.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

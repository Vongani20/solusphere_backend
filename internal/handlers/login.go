package handlers

import (
	"Solusphere/internal/db"
	"database/sql"
	"encoding/json"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	json.NewDecoder(r.Body).Decode(&req)

	var id int
	var password string
	err := db.DB.QueryRow("SELECT id, password FROM users WHERE email=?", req.Email).Scan(&id, &password)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	// TODO: Compare password (hash check)
	w.Write([]byte("Login successful"))
}

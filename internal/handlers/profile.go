package handlers

import (
	"net/http"
)

// Profile handler (must be exported, starts with uppercase)
func Profile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is the profile page"))
}

// api/login.go

package handler

import (
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// In a real-world scenario, you'd validate credentials and set a proper session/token.
		if username == "admin" && password == "password" {
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
			return
		}
	}

	http.Error(w, "Invalid login credentials", http.StatusUnauthorized)
}

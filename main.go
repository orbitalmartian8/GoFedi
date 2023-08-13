package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var (
	tmpl *template.Template
)

func init() {
	tmpl = template.Must(template.ParseFiles("login.html"))
}

func main() {
	http.HandleFunc("/", loginHandler)
	http.HandleFunc("/dashboard", dashboardHandler)
	port := ":8080"
	fmt.Printf("Server listening on port %s...\n", port)
	http.ListenAndServe(port, nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// In a real-world scenario, you'd validate credentials and set a proper session/token.
		if username == "admin" && password == "password" {
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
			return
		}
	}

	tmpl.Execute(w, nil)
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the dashboard!")
}

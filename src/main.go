// main.go
package main

import (
    "net/http"
    "html/template"
    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()
    
    config := LoadConfig()
    
    // Define routes and handlers
    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl := template.Must(template.ParseFiles("templates/index.html"))
        tmpl.Execute(w, nil)
    })
    
    router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
        instance := r.FormValue("instance")
        userId := r.FormValue("userId")
        
        // Save instance and userId to session or database
        
        http.Redirect(w, r, "/timeline", http.StatusSeeOther)
    })
    
    router.HandleFunc("/timeline", func(w http.ResponseWriter, r *http.Request) {
        tmpl := template.Must(template.ParseFiles("templates/timeline.html"))
        tmpl.Execute(w, nil)
    })
    
    // Initialize server
    http.Handle("/", router)
    http.ListenAndServe(":8080", nil)
}

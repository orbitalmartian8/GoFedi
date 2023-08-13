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
    
    // Initialize server
    http.Handle("/", router)
    http.ListenAndServe(":8080", nil)
}

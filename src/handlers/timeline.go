// handlers/timeline.go
package handlers

import (
    "net/http"
    "html/template"
    "github.com/gorilla/mux"
)

func TimelineHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/timeline.html"))
    tmpl.Execute(w, nil)
}

package main

import (
  "net/http"
  "gigs/controllers"
  "gigs/models"
  "github.com/gorilla/mux"
)

func main() {
  // Load DB in models package
  models.LoadDB("postgres://www:www@localhost/gigs")

  // Instantiate gorilla mux
  r := mux.NewRouter()

  // Define routes
  r.HandleFunc("/", controllers.IndexHandler)
  r.HandleFunc("/gigs", controllers.GigsHandler)
  r.HandleFunc("/gigs/{id:[0-9]+}", controllers.GigHandler)

  // Start server with gorilla mux
  http.Handle("/", r)
  http.ListenAndServe(":8080", r)
}


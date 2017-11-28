package main

import (
  "net/http"
  "gigs/controllers"
  "gigs/models"
)

func main() {
  // Load DB in models package
  models.LoadDB("postgres://www:www@localhost/gigs")

  // Define routes
  http.HandleFunc("/", controllers.IndexHandler)
  http.HandleFunc("/gigs", controllers.GigsHandler)

  // Start server with default mux
  http.ListenAndServe(":8080", nil)
}


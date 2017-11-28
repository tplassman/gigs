package controllers

import (
  "html/template"
  "net/http"
  "gigs/models"
)

func GigsHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method != "GET" {
    http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
    return
  }

  gigs, err := models.AllGigs()
  if err != nil {
    http.Error(w, http.StatusText(500), 500)
    return
  }

  t, err := template.ParseFiles("views/layout.tmpl", "views/gigs.tmpl")
  if err != nil {
    panic(err)
  }
  t.Execute(w, gigs)
}


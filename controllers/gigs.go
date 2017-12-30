package controllers

import (
	"gigs/models"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
)

func GigsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	gigs, err := models.GetGigs()
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

func GigHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	gig, err := models.GetGig(uint(id))
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	t, err := template.ParseFiles("views/layout.tmpl", "views/gig.tmpl")
	if err != nil {
		panic(err)
	}
	t.Execute(w, gig)
}

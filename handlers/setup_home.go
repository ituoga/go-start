package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
)

func SetupHome(router chi.Router, session sessions.Store) {

	homeHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Home"))
	}

	router.Get("/", homeHandler)
}

package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *application) routes() *chi.Mux {
	// application routes
	r := chi.NewRouter()

	// add middleware

	// add your routes here
	r.Get("/", a.Handlers.Home) // default home route

	// static routes
	fileServer := http.FileServer(http.Dir("./public"))
	r.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return r
}

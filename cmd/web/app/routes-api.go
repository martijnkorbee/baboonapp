package app

import (
	"github.com/go-chi/chi/v5"
)

func (a *application) routesAPI() *chi.Mux {
	// API routes
	r := chi.NewRouter()

	// add your middleware here

	// add your routes here
	r.Get("/ping", a.Handlers.Ping) // default route

	return r
}

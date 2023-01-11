package routes

import (
	"github.com/go-chi/chi/v5"
)

func (ar *AppRoutes) routesAPI() *chi.Mux {
	// API routes
	r := chi.NewRouter()

	// add your middleware here

	// add your routes here
	r.Get("/ping", a.Handlers.Ping) // default route

	return r
}

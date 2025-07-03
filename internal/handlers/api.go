package handlers

import (
	"github.com/go-chi/chi/v5"
	chimiddle "github.com/go-chi/chi/v5/middleware"
	"github.com/renvins/gocoins_restful/internal/middleware"
)

// With first letter capitalized to export the function
func Handler(r *chi.Mux) {
	// Middleware is a function called before the primary handler
	// It's used to route requests, log them, or handle errors
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)
	})
}

package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()

	// register standard middlewares
	r.Use(app.recoverPanic)
	r.Use(app.logRequest)
	r.Use(secureHeader)

	r.Get("/ping", ping)

	// register handler functions to routes
	r.Group(func(r chi.Router) {
		r.Use(app.session.Enable)
		r.Use(noSurf)
		r.Use(app.authenticate)

		r.Get("/", app.home)

		r.Route("/snippet", func(r chi.Router) {
			r.Get("/{id}", app.showSnippet)

			r.Group(func(r chi.Router) {
				r.Use(app.requireAuthenticatedUser)
				r.Get("/create", app.createSnippetForm)
				r.Post("/create", app.createSnippet)
			})
		})
		r.Route("/user", func(r chi.Router) {
			r.Get("/signup", app.signUpFormUser)
			r.Post("/signup", app.signUpUser)
			r.Get("/login", app.loginFormUser)
			r.Post("/login", app.loginUser)

			r.Group(func(r chi.Router) {
				r.Use(app.requireAuthenticatedUser)
				r.Post("/logout", app.logoutUser)
			})
		})
	})
	// create route/handler to serve static assets
	fileServer := http.FileServer(http.Dir("./ui/static"))
	r.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return r
}

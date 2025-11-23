package main

import (
	"log"
	"net/http"
	"test/internal/store"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
	store  store.Storage
}

type config struct {
	addr string
}

// Function to mount routes
func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	// Use Route to mount sub-routes with "/v1"
	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
	})

	return r
}

// Function to start the server
func (app *application) start(mux http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30, // The maximum duration for writing a response.
		ReadTimeout:  time.Second * 10, // The maximum duration for reading a request.
		IdleTimeout:  time.Minute,      // The maximum duration for an idle connection.
	}

	log.Printf("The server has started at %s", app.config.addr)

	return srv.ListenAndServe()
}

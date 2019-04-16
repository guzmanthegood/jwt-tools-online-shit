package main

import (
	"log"
	"net/http"

	"jwt-tools-online-shit/auth"
	"jwt-tools-online-shit/config"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func newRouter(c config.AppConfig) *chi.Mux {
	r := chi.NewRouter()

	// set response type as json in all server
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// recover from panics
	r.Use(Recoverer)

	// log transactions
	if c.IsDevelopment() {
		r.Use(middleware.Logger)
	}

	// api routes
	r.Route("/api/v1", func(r chi.Router) {
		r.Mount("/auth", auth.Routes(c))
		r.Get("/status", Status)
		r.Get("/panic", Panic)
		r.NotFound(NotFound)
	})

	// log all routes
	walkFunc := func(method string, route string, handlder http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("[INFO] %s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(r, walkFunc); err != nil {
		log.Panicf("[ERRO] logging error: %s\n", err.Error())
	}

	return r
}

package main

import (
	"log"
	"net/http"

	"jwt-tools-online-shit/utils"

	"github.com/go-chi/render"
)

// NotFound render status function
func NotFound(w http.ResponseWriter, r *http.Request) {
	log.Printf("[ERRO] resource not found: %s\n", r.URL.Path)
	render.Render(w, r, utils.NotFound(r.URL.Path))
}

// Status function
func Status(w http.ResponseWriter, r *http.Request) {
	type Status struct {
		Ok bool `json:"ok"`
	}
	ok := Status{Ok: true}
	render.JSON(w, r, ok)
}

// Panic status function
func Panic(w http.ResponseWriter, r *http.Request) {
	panic("panic test")
}

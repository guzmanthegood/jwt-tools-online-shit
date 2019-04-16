package main

import (
	"errors"
	"log"
	"net/http"

	"jwt-tools-online-shit/utils"

	"github.com/go-chi/render"
)

// Recoverer middleware to rescue from panic
func Recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				switch x := err.(type) {
				case string:
					err = errors.New(x)
				case error:
					err = x
				default:
					err = errors.New("unknown panic")
				}

				log.Printf("[ERRO] internal server error: %v\n", err)
				render.JSON(w, r, utils.InternalError(err.(error)))
			}

		}()
		next.ServeHTTP(w, r)
	})
}

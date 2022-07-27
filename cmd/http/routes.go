package main

import (
	"flights-assignment/cmd/http/handler"
	"flights-assignment/internal/decoder"
	"flights-assignment/internal/marshaler"
	"flights-assignment/internal/reswriter"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func routes() chi.Router {
	r := chi.NewRouter()

	/*
	** I'd like to use the clean path middleware to remove additional slashes
	** but there is an issue between the go version and the library on my machine.
	**
	** The program works correctly when the middleware is not in use.
	 */
	//r.Use(middleware.CleanPath)
	r.MethodFunc(http.MethodPost, "/track",
		handler.PostTrack(
			reswriter.New(marshaler.New()), decoder.New(),
		),
	)
	r.MethodFunc(http.MethodGet, "/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		log.Printf("log level 'trace' - Health: responded successfully with status code '%d'",
			http.StatusOK)
	})
	return r
}

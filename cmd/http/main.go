package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	err := run()
	if err != nil {
		log.Fatalf("log level 'fatal' - main: '%s'", err)
	}
}

func run() error {
	port := "8080"
	log.Printf("log level 'info' - main: starting application on port '%s'...", port)
	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           routes(),
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		IdleTimeout:       10 * time.Second,
	}
	err := srv.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

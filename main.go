package main

//go:generate swagger generate spec

// Package classification Pokemon API.
//
// the purpose of this application is to provide an application
// that is using plain go code and the latest version of net/http to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http
//     Host: localhost
//     BasePath: /
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - api_key:
//
// swagger:meta

import (
	"fmt"
	"ketchum/cfg"
	"ketchum/items"
	"ketchum/middleware"
	"ketchum/pokemon"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	fmt.Println("Server starting...")

	// Create a new ServeMux
	mux := http.NewServeMux()
	db := cfg.InitConfig()
	// want to add a logger something like zap.Logger

	pRepo := pokemon.NewPokemonRepo(db)
	iRepo := items.NewItemRepo(db)
	pService := pokemon.NewPokemonService(pRepo)
	iService := items.NewItemService(iRepo)
	pService.RegisterHandlers(mux)
	iService.RegisterHandlers(mux)

	// Set up middleware for logging requests
	loggedHandler := middleware.LogMiddleware(mux)

	// ideally these could be set in the config via env variables
	// Set up rate limiting
	limiter := rate.NewLimiter(rate.Limit(25), 1) // Allow 25 requests per second
	// Set up timeout handler
	timeout := 1 * time.Minute

	// Wrap the loggedHandler with rate limiting and timeout
	rateLimitedHandler := middleware.RateLimiterMiddleware(loggedHandler, limiter)
	timeoutHandler := http.TimeoutHandler(rateLimitedHandler, timeout, "Request timed out")

	go func() {
		if err := http.ListenAndServe(":8080", timeoutHandler); err != nil {
			log.Fatal(err)
		}
	}()
	signs := make(chan os.Signal, 1)
	signal.Notify(signs, syscall.SIGINT, syscall.SIGTERM)
	<-signs
	// Gracefully shutdown the server
	log.Println("Shutting down...")
	log.Println("Server gracefully stopped")
}

package probot

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	portVar int
	ipVar   string
)

var app *App

// Router creates a new mux.Router and registers our webhook handler
func Router(path string) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc(path, rootHandler(app)).Methods("POST")

	return r
}

// Start handles initialization and setup of the webhook server
func Start() {
	initialize()

	// Webhook router
	router := Router("/")

	// Server
	log.Printf("Server running at: http://localhost:%d/\n", portVar)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", ipVar, portVar), router))
}

func initialize() {
	// Parse incoming command-line arguments
	flag.IntVar(&portVar, "p", 8000, "port to listen on (default: 8000)")
	flag.StringVar(&ipVar, "i", "127.0.0.1", "interface to listen on (default: 127.0.0.1)")
	flag.Parse()

	// Initialize app
	app = NewApp()
	log.Printf("Loaded GitHub App ID: %d\n", app.ID)
}

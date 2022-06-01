package probot

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	portVar int
	ipVar   string
)

var app *App

// Start handles initialization and setup of the webhook server
func Start() {
	initialize()

	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler(app))

	// Server
	log.Printf("Server running at: http://%s:%d/\n", ipVar, portVar)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", ipVar, portVar), mux))
}

func StartArgs(iface string, port string) {
	initialize()
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler(app))

	// Server
	log.Printf("Server running at: http://%s:%d/\n", iface, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", iifacepVar, port), mux))
}

func initialize() {
	// Parse incoming command-line arguments
	flag.IntVar(&portVar, "p", 8000, "port to listen on (default: 8000)")
	flag.StringVar(&ipVar, "i", "127.0.0.1", "IP address to listen on (default: 127.0.0.1)")
	flag.Parse()

	// Initialize app
	app = NewApp()
	log.Printf("Loaded GitHub App ID: %d\n", app.ID)
}

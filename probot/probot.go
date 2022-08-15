package probot

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

var (
	portVar int
	ipVar   string
	log     *zap.SugaredLogger
)

var app *App

// Start handles initialization and setup of the webhook server
func Start() {
	log = newLogger()
	StartArgs("0.0.0.0", 8080, 9999)
}

func StartArgs(iface string, port int, healthPort int) {
	initialize()

	// Set up health check
	healthMux := http.NewServeMux()
	healthMux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})
	go func() {
		http.ListenAndServe(fmt.Sprintf("%s:%d", iface, healthPort), healthMux)
	}()

	// Set up server
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler(app))
	log.Infof("Server running at: http://%s:%d/\n", iface, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", iface, port), mux))
}

func initialize() {
	// Initialize app
	app = NewApp()
	log.Infof("Loaded GitHub App ID: %d\n", app.ID)
}

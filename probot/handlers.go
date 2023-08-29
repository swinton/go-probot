package probot

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/go-github/v53/github"
)

var handlers = make(map[string]eventHandler)

func rootHandler(app *App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		context := NewContext(app)

		// Validate the payload
		// Per the docs: https://docs.github.com/en/developers/webhooks-and-events/securing-your-webhooks#validating-payloads-from-github
		payloadBytes, err := github.ValidatePayload(r, []byte(app.Secret))
		if err != nil {
			log.Println(err)
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		// Log the request headers
		log.Printf("Signature validates: %s\n", r.Header.Get("X-Hub-Signature"))
		log.Printf("Headers: %v\n", r.Header)

		// Get the installation from the payload
		payload := &payloadInstallation{}
		json.Unmarshal(payloadBytes, payload)
		log.Printf("Installation: %d\n", payload.Installation.GetID())
		log.Printf("Received GitHub App ID %d\n", app.ID)

		// Parse the incoming request into an event
		context.Payload, err = github.ParseWebHook(github.WebHookType(r), payloadBytes)
		if err != nil {
			log.Println(err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		log.Printf("Event type: %T\n", context.Payload)

		// Instantiate client
		installation := &installation{ID: payload.Installation.GetID()}
		context.GitHub, err = NewEnterpriseClient(app, installation)
		if err != nil {
			log.Println(err)
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}
		log.Printf("client %s instantiated for %s\n", context.GitHub.UserAgent, context.GitHub.BaseURL)

		// Reset the body for subsequent handlers to access
		r.Body = reset(r.Body, payloadBytes)

		// Look to see if we have a handler for the incoming webhook type
		if handler, ok := handlers[github.WebHookType(r)]; ok {
			err = handler(context)
			if err != nil {
				log.Println(err)
				http.Error(w, "Server Error", http.StatusInternalServerError)
				return
			}
		} else {
			log.Printf("Unknown event type: %s\n", github.WebHookType(r))
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		// Success!
		// Send response as application/json
		resp := webhookResponse{
			Received: true,
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}

// HandleEvent registers an eventHandler for a named event
func HandleEvent(event string, f eventHandler) {
	handlers[event] = f
}

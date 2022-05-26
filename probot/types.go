package probot

import "github.com/google/go-github/v45/github"

type eventHandler func(ctx *Context) error

type webhookResponse struct {
	Received bool `json:"received"`
}

// payloadInstallation represents the incoming installation part of the payload
type payloadInstallation struct {
	Installation *github.Installation `json:"installation"`
}

// Installation encapsulates the fields needed to define an installation of a GitHub App
type installation struct {
	ID int64
}

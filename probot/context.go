package probot

import "github.com/google/go-github/github"

// Context encapsulates the fields passed to webhook handlers
type Context struct {
	App     *App
	Payload interface{}
	GitHub  *github.Client
}

// NewContext instantiates a new context
func NewContext(app *App) *Context {
	return &Context{App: app}
}

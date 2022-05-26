package main

import (
	"context"
	"log"

	"github.com/ccureau/go-probot/probot"
	"github.com/google/go-github/v45/github"
)

func main() {
	// Add a handler for "issues" events
	probot.HandleEvent("issues", func(ctx *probot.Context) error {
		// Because we're listening for "issues" we know the payload is a *github.IssuesEvent
		event := ctx.Payload.(*github.IssuesEvent)
		log.Printf("🌈 Got issues %+v\n", event)

		// Create a comment back on the issue
		// https://github.com/google/go-github/blob/d57a3a84ba041135efb6b7ad3991f827c93c306a/github/issues_comments.go#L101-L117
		newComment := &github.IssueComment{Body: github.String("## :wave: :earth_americas:\n\n![fellowshipoftheclaps](https://user-images.githubusercontent.com/27806/91333726-91c46f00-e793-11ea-9724-dc2e18ca28d0.gif)")}
		comment, _, err := ctx.GitHub.Issues.CreateComment(context.Background(), *event.Repo.Owner.Login, *event.Repo.Name, int(event.Issue.GetID()), newComment)
		if err != nil {
			return err
		}

		// Log out our new comment
		log.Printf("✨ New comment created: %+v\n", comment)

		return nil
	})

	probot.Start()
}

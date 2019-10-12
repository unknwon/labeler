package main

import (
	"context"

	"github.com/google/go-github/v28/github"
	"github.com/urfave/cli"
	"golang.org/x/oauth2"
)

// client returns a new GitHub v4 client.
func client(c *cli.Context) *github.Client {
	return github.NewClient(
		oauth2.NewClient(
			context.Background(),
			oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: c.GlobalString("token")},
			),
		),
	)
}

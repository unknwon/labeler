package main

import (
	"context"

	"github.com/shurcooL/githubv4"
	"github.com/urfave/cli"
	"golang.org/x/oauth2"
)

// client returns a new GitHub v4 client.
func client(c *cli.Context) *githubv4.Client {
	return githubv4.NewClient(
		oauth2.NewClient(
			context.Background(),
			oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: c.GlobalString("token")},
			),
		),
	)
}

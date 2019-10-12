package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/go-github/v28/github"
	"github.com/urfave/cli"
)

var syncCmd = cli.Command{
	Name:        "sync",
	Usage:       "Sync labels from a template file to target repository",
	Description: `labeler sync --owner=unknwon --repo=labeler --from unknwon_labeler.json`,
	Action:      runSyncCmd,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "owner",
			Usage:    "GitHub owner handle (e.g. unknwon)",
			Required: true,
		},
		cli.StringFlag{
			Name:     "repo",
			Usage:    "GitHub repository handle (e.g. labeler)",
			Required: true,
		},
		cli.StringFlag{
			Name:     "from",
			Usage:    "Template file path (e.g. unknwon_labeler.json)",
			Required: true,
		},
	},
}

func runSyncCmd(c *cli.Context) error {
	fr, err := os.Open(c.String("from"))
	if err != nil {
		return fmt.Errorf("open file: %v", err)
	}
	defer fr.Close()

	var labels []*Label
	err = json.NewDecoder(fr).Decode(&labels)
	if err != nil {
		return fmt.Errorf("decode JSON: %v", err)
	}

	client := client(c)

createLabel:
	for _, l := range labels {
		_, _, err = client.Issues.CreateLabel(context.Background(), c.String("owner"), c.String("repo"), &github.Label{
			Name:        github.String(l.Name),
			Color:       github.String(l.Color),
			Description: github.String(l.Description),
		})
		if er, ok := err.(*github.ErrorResponse); ok {
			for _, e := range er.Errors {
				if e.Field == "name" && e.Code == "already_exists" {
					continue createLabel
				}
			}
		}
		if err != nil {
			return fmt.Errorf("create label %q: %v", l.Name, err)
		}
	}

	fmt.Printf("Labels in %q have been synced to '%s/%s'!\n", c.String("from"), c.String("owner"), c.String("repo"))
	return nil
}

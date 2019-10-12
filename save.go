package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/go-github/v28/github"
	"github.com/urfave/cli"
)

var saveCmd = cli.Command{
	Name:        "save",
	Usage:       "Save labels of target repository to a template file",
	Description: `labeler save --owner=unknwon --repo=labeler --to unknwon_labeler.json`,
	Action:      runSaveCmd,
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
			Name:     "to",
			Usage:    "Template file path (e.g. unknwon_labeler.json)",
			Required: true,
		},
	},
}

func runSaveCmd(c *cli.Context) error {
	opt := &github.ListOptions{
		PerPage: 100,
	}
	client := client(c)

	var labels []*Label
	for {
		ls, resp, err := client.Issues.ListLabels(context.Background(), c.String("owner"), c.String("repo"), opt)
		if err != nil {
			return fmt.Errorf("list labels: %v", err)
		}
		for _, l := range ls {
			labels = append(labels, ToLabel(l))
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	fw, err := os.Create(c.String("to"))
	if err != nil {
		return fmt.Errorf("create file: %v", err)
	}
	defer fw.Close()

	enc := json.NewEncoder(fw)
	enc.SetIndent("", "  ")
	if err = enc.Encode(labels); err != nil {
		return fmt.Errorf("encode JSON: %v", err)
	}
	return nil
}

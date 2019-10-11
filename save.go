package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/shurcooL/githubv4"
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
	client := client(c)

	var q struct {
		Repository struct {
			Labels struct {
				Nodes    []Label
				PageInfo PageInfo
			} `graphql:"labels(first: 100, after: $after)"`
		} `graphql:"repository(owner: $repositoryOwner,name: $repositoryName)"`
	}
	vars := map[string]interface{}{
		"repositoryOwner": githubv4.String(c.String("owner")),
		"repositoryName":  githubv4.String(c.String("repo")),
		"after":           (*githubv4.String)(nil),
	}

	var labels []Label
	for {
		err := client.Query(context.Background(), &q, vars)
		if err != nil {
			return fmt.Errorf("query: %v", err)
		}

		labels = append(labels, q.Repository.Labels.Nodes...)

		if !q.Repository.Labels.PageInfo.HasNextPage {
			break
		}

		vars["after"] = githubv4.String(q.Repository.Labels.PageInfo.EndCursor)
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

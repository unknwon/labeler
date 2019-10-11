package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

const version = "0.1.0"

func main() {
	log.SetFlags(0)

	app := &cli.App{
		Name:  "labeler",
		Usage: "A CLI tool to sync labels for a GitHub repository with templates",
		Commands: []cli.Command{
			saveCmd,
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:     "token",
				Usage:    "GitHub personal access token",
				EnvVar:   "LABELER_TOKEN",
				Required: true,
			},
		},
		Version: version,
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

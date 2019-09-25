package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

const version = "0.0.1"

func main() {
	log.SetFlags(0)

	(&cli.App{
		Name:     "labeler",
		Usage:    "A CLI tool to sync labels for a GitHub repository with templates",
		Commands: nil,
		Flags:    nil,
		Version:  version,
	}).Run(os.Args)
}

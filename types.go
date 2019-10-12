package main

import (
	"github.com/google/go-github/v28/github"
)

// Label contains information of a GitHub label.
type Label struct {
	Name        string
	Color       string
	Description string
}

// ToLabel populates non-nil value from github.Label to Label.
func ToLabel(l *github.Label) *Label {
	var label Label
	if l.Name != nil {
		label.Name = *l.Name
	}
	if l.Color != nil {
		label.Color = *l.Color
	}
	if l.Description != nil {
		label.Description = *l.Description
	}
	return &label
}

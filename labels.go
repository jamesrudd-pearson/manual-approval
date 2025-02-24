package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/v43/github"
)

// createLabelIfNotExists creates a label if it does not exist.
// It returns an error if the label does not exist and it fails to create it.
// It returns nil if the label already exists or if it was successfully created.
func createLabelIfNotExists(client *github.Client, repoFullName string, label string) error {
	repoOwner := strings.Split(repoFullName, "/")[0]
	repoName := strings.Split(repoFullName, "/")[1]

	_, resp, err := client.Issues.GetLabel(context.Background(), repoOwner, repoName, label)
	if err != nil {
		if resp.StatusCode != 404 {
			return fmt.Errorf("error getting label: %w", err)
		}
	}

	if resp.StatusCode == 200 {
		return nil
	}

	fmt.Printf("Label \"%s\" does not exist in \"%s/%s\". Creating it...\n", label, repoOwner, repoName)

	_, _, err = client.Issues.CreateLabel(context.Background(), repoOwner, repoName, &github.Label{
		Name: &label,
	})
	if err != nil {
		return fmt.Errorf("error creating label: %w", err)
	}

	return nil
}

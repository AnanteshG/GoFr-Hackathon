package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
)

const githubToken = "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"

// GetClient returns an authenticated GitHub client
func GetClient() *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(context.Background(), ts)
	return github.NewClient(tc)
}

// VerifyRepository checks if the repository exists and is accessible
func VerifyRepository(client *github.Client, owner, repo string) error {
	repository, resp, err := client.Repositories.Get(context.Background(), owner, repo)
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			return fmt.Errorf("repository %s/%s not found. Please check the repository name and your access permissions", owner, repo)
		}
		return fmt.Errorf("error verifying repository: %v", err)
	}

	fmt.Printf("Successfully found repository: %s\n", *repository.FullName)
	if *repository.Private {
		fmt.Println("Note: This is a private repository")
	}
	return nil
}

// SaveToDataFile appends content to the data.txt file
func SaveToDataFile(content string) error {
	file, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open data.txt: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(content + "\n\n")
	if err != nil {
		return fmt.Errorf("failed to write to data.txt: %v", err)
	}

	return nil
}

// FetchRepoContent recursively fetches the file structure and code of a GitHub repo
func FetchRepoContent(client *github.Client, owner, repo, path string) error {
	_, directoryContent, resp, err := client.Repositories.GetContents(context.Background(), owner, repo, path, nil)
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			return fmt.Errorf("path '%s' not found in repository", path)
		}
		return fmt.Errorf("error fetching contents at path '%s': %v", path, err)
	}

	// Process all items in the directory
	for _, content := range directoryContent {
		if *content.Type == "file" {
			// Download and save file content
			fileContent, _, err := client.Repositories.DownloadContents(context.Background(), owner, repo, *content.Path, nil)
			if err != nil {
				return fmt.Errorf("failed to download file %s: %v", *content.Path, err)
			}

			contentBytes, err := io.ReadAll(fileContent)
			if err != nil {
				return fmt.Errorf("failed to read file %s: %v", *content.Path, err)
			}

			err = SaveToDataFile(fmt.Sprintf("File: %s\n%s", *content.Path, string(contentBytes)))
			if err != nil {
				return fmt.Errorf("failed to save content of %s: %v", *content.Path, err)
			}
			fmt.Printf("Processed file: %s\n", *content.Path)
		} else if *content.Type == "dir" {
			// Recursively process subdirectories
			fmt.Printf("Entering directory: %s\n", *content.Path)
			err := FetchRepoContent(client, owner, repo, *content.Path)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// This code returns all the contents of files and Directories in a GitHub Repository.

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
)

// GetClient returns an authenticated GitHub client
func GetClient(token string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
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

// FetchRepoContent recursively fetches the file structure and code of a GitHub repo
func FetchRepoContent(client *github.Client, owner, repo, path string) ([]*github.RepositoryContent, error) {
	var allContents []*github.RepositoryContent

	// Fetch contents at the current path
	contents, dirContents, _, err := client.Repositories.GetContents(context.Background(), owner, repo, path, nil)
	if err != nil {
		return nil, fmt.Errorf("error fetching contents at path '%s': %v", path, err)
	}

	// Handle single file case
	if contents != nil && dirContents == nil {
		allContents = append(allContents, contents)
		return allContents, nil
	}

	// Handle directory case
	for _, content := range dirContents {
		if *content.Type == "dir" {
			fmt.Println("Directory:", *content.Path)
			subContents, err := FetchRepoContent(client, owner, repo, *content.Path)
			if err != nil {
				return nil, err
			}
			allContents = append(allContents, subContents...)
		} else if *content.Type == "file" {
			fmt.Println("File:", *content.Path)
			allContents = append(allContents, content)
		}
	}

	return allContents, nil
}

func main() {
	// Replace with your actual GitHub token
	token := "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"

	// Updated repository name with correct case
	owner := "AnanteshG"
	repo := "Contest-Hub"

	client := GetClient(token)

	// First verify the repository exists and is accessible
	err := VerifyRepository(client, owner, repo)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Starting to fetch repository content...")

	// Fetch repository content
	_, err = FetchRepoContent(client, owner, repo, "")
	if err != nil {
		log.Fatal("Error fetching repository content: ", err)
	}

	fmt.Println("Repository scraping completed!")
}

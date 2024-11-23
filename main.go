package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	

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

// SaveFile saves a file's content to a local file
func SaveFile(path, content string) error {
	// Clean the path to handle any potential directory traversal
	cleanPath := filepath.Clean(path)

	// Create the necessary directories for the path if they don't exist
	dir := filepath.Dir(cleanPath)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directory %s: %v", dir, err)
	}

	// If the path already exists and is a directory, return an error
	fileInfo, err := os.Stat(cleanPath)
	if err == nil && fileInfo.IsDir() {
		return fmt.Errorf("cannot save file: %s is a directory", cleanPath)
	}

	// Write the file content
	err = ioutil.WriteFile(cleanPath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("failed to write file %s: %v", cleanPath, err)
	}

	fmt.Printf("Saved file: %s\n", cleanPath)
	return nil
}



// FetchRepoContent recursively fetches the file structure and code of a GitHub repo
func FetchRepoContent(client *github.Client, owner, repo, path string) ([]*github.RepositoryContent, error) {
	var allContents []*github.RepositoryContent

	// Fetch repository contents for the current path
	contents, dirContents, _, err := client.Repositories.GetContents(context.Background(), owner, repo, path, nil)
	if err != nil {
		return nil, fmt.Errorf("error fetching contents at path '%s': %v", path, err)
	}

	// Handle if it's a file and not a directory
	if contents != nil && dirContents == nil {
		if *contents.Type == "file" {
			fileContent, _, err := client.Repositories.DownloadContents(context.Background(), owner, repo, path, nil)
			if err != nil {
				return nil, fmt.Errorf("failed to download %s: %v", path, err)
			}
			contentBytes, err := ioutil.ReadAll(fileContent)
			if err != nil {
				return nil, fmt.Errorf("failed to read %s: %v", path, err)
			}
			err = SaveFile(path, string(contentBytes))
			if err != nil {
				return nil, err
			}
			allContents = append(allContents, contents)
		}
		return allContents, nil
	}

	// Handle directory case
	for _, content := range dirContents {
		if *content.Type == "dir" {
			// Recursively fetch contents from subdirectories
			subContents, err := FetchRepoContent(client, owner, repo, *content.Path)
			if err != nil {
				return nil, err
			}
			allContents = append(allContents, subContents...)
		} else if *content.Type == "file" {
			// Handle file content
			fileContent, _, err := client.Repositories.DownloadContents(context.Background(), owner, repo, *content.Path, nil)
			if err != nil {
				return nil, fmt.Errorf("failed to download %s: %v", *content.Path, err)
			}
			contentBytes, err := ioutil.ReadAll(fileContent)
			if err != nil {
				return nil, fmt.Errorf("failed to read %s: %v", *content.Path, err)
			}
			err = SaveFile(*content.Path, string(contentBytes))
			if err != nil {
				return nil, err
			}
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
	repo := "Contest-Hub"  // Fixed repository name
	
	client := GetClient(token)

	// First verify the repository exists and is accessible
	err := VerifyRepository(client, owner, repo)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Starting to fetch repository content...")
	
	_, err = FetchRepoContent(client, owner, repo, "")
	if err != nil {
		log.Fatal("Error fetching repository content: ", err)
	}
	
	fmt.Println("Repository scraping completed!")
}
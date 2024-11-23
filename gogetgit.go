package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
)

func main() {
	// Ensure correct usage
	if len(os.Args) < 2 {
		fmt.Println("Usage: gogetgit <GitHub URL>")
		os.Exit(1)
	}

	// Parse the GitHub URL
	githubURL := os.Args[1]
	parsedURL, err := url.Parse(githubURL)
	if err != nil || parsedURL.Host != "github.com" {
		log.Fatalf("Invalid GitHub URL: %s", githubURL)
	}

	// Extract owner and repository from the URL
	parts := strings.Split(strings.Trim(parsedURL.Path, "/"), "/")
	if len(parts) < 2 {
		log.Fatalf("Invalid GitHub URL format. Expected: https://github.com/<owner>/<repo>")
	}
	owner, repo := parts[0], parts[1]

	// Clear existing data.txt file
	err = os.Remove("data.txt")
	if err != nil && !os.IsNotExist(err) {
		log.Fatalf("Failed to clear data.txt: %v", err)
	}

	// Authenticate and fetch repository content
	client := GetClient()
	err = VerifyRepository(client, owner, repo)
	if err != nil {
		log.Fatalf("Error verifying repository: %v", err)
	}

	fmt.Println("Starting to scrape repository...")
	err = FetchRepoContent(client, owner, repo, "")
	if err != nil {
		log.Fatalf("Error fetching repository content: %v", err)
	}

	fmt.Println("Repository scraping completed! Check data.txt for details.")
}

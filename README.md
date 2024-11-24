
# GoGetGit

GoGetGit is a CLI tool built using the GoFr framework for creating API endpoints. The project scrapes code from a given GitHub repository using the Go-GitHub package, stores code in a vector format using ChromaDB, and processes data with LLaMA 3.2 for natural language understanding. The tool can be run from a terminal with the command `gogetgit <GitHub_url>`.

---

## Disclaimer
 ![#f03c15](https://placehold.co/15x15/f03c15/f03c15.png)
**This is a prototype built in a 24-hr hackathon. The web version will be released shortly and the gogetgit package will be published online to directly download it into your PCs. Many features yet to come...**.
Feel Free to contribute.

---

## Features

- **GoFr Framework**: Used for building endpoints that handle GitHub repository scraping and data processing.
- **Go-GitHub**: A Go client for GitHub APIs, used for scraping the entire code from a given GitHub repository.
- **ChromaDB**: A vector storage system used to store and index the code data for quick retrieval.
- **LLaMA 3.2**: A language model used for processing and understanding the scraped code.
- **CLI Tool**: A simple terminal command `gogetgit <GitHub_url>` to initiate the scraping and processing of a GitHub repository.

## Requirements

- Go 1.18 or higher
- GoFr framework
- Go-GitHub package
- ChromaDB
- LLaMA 3.2 model (can be integrated via an API or local setup)
- GitHub account (to access repositories)

## Installation

1. Clone the repository:
 ```bash
    git clone https://github.com/AnanteshG/GoFr-Hackathon.git
   ``` 
2. Run the Backend LLAMA Server:
 ```bash
    python ollama.py
   ``` 
3. Install dependencies:
 ```bash
    go mod tidy
   ``` 
   4. Run the main .go script:
   ```bash
    go run gogetgit.go <GITHUB_URL>
   ```
    
## Usage

To use the CLI tool, run the following command in your terminal:

```bash
gogetgit <GitHub_url>
```

Replace `<GitHub_url>` with the URL of the GitHub repository you want to scrape. The tool will:

1.  Clone the repository.
2.  Extract all code files.
3.  Store code in ChromaDB as vectors.
4.  Use LLaMA 3.2 to process the code data (optional, for advanced processing).
 

## How It Works

-   **GitHub Scraping**: The `go-github` package is used to fetch all files from the specified GitHub repository.
-   **Vector Storage**: Code files are processed into vectors and stored in ChromaDB. This enables fast search and retrieval of code snippets.
-   **Data Processing with LLaMA 3.2**: Once the code is scraped and stored, LLaMA 3.2 is used to analyze and understand the content, allowing for advanced querying or processing based on the repository's codebase.

## Contributing

Feel free to fork the repository, make changes, and submit pull requests. Contributions are welcome!



 ### Explanation:
- **Project Overview**: Describes the overall functionality and components.
- **Installation Instructions**: Walks through cloning the repo, navigating to it, and installing dependencies.
- **Usage**: Provides a basic example of how to run the CLI tool.
- **How It Works**: Explains the internal process, such as using `go-github` for scraping, `ChromaDB` for storage, and `LLaMA 3.2` for data processing.
- **Contributing & License**: Common sections that encourage contributions and clarify the license.

This template can be customized further based on any additional details you might want to include!

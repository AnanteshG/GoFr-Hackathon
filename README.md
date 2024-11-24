# GitHub Repository Chat App

## Overview

This web application allows users to input a GitHub repository URL, scrape its data, and save the extracted information as a dataset in **LLAMA 3.2**. Users can then chat with the repository, asking about its contents, structure, and other details.

The project is built using **Go**, the **GoFr framework**, and **Flask**, ensuring a scalable and high-performance application while offering flexibility for AI interaction.

---

## Features

- **GitHub Repository Scraper**: Extract metadata, files, and commit history from any public GitHub repository.
- **Dataset Creation**: Save the scraped data in a format compatible with LLAMA 3.2 for efficient querying.
- **Interactive Chat**: Use natural language to inquire about the repository's contents and structure.
- **Fast and Scalable**: Built with Go and GoFr for optimal performance.
- **CLI Tool**: Automates the scraping process for GitHub repositories, providing a convenient command-line interface.
- **Flask Webserver**: Facilitates interaction with LLAMA 3.2, enabling users to send queries and receive responses seamlessly.

---

## Tech Stack

### Backend
- **Go**: Main programming language.
- **GoFr Framework**: To structure and manage the web application.
- **GitHub API**: To fetch repository data.
- **go-github**: A Go library for interacting with the GitHub API.

### AI Integration
- **LLAMA 3.2**: Used for dataset processing and enabling interactive chat.

### Additional Tools
- **Flask**: Used to create a webserver for handling user requests and enabling chat functionality.
- **CLI Tool(gogetgit.go)**: Built to automate the GitHub scraping process.
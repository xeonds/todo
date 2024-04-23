package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"todo/lib"
)

type Config struct {
	serverUrl string
}

func main() {
	config := lib.LoadConfig[Config]()
	if len(os.Args) < 2 {
		fmt.Println("Usage: client [push|pull] [filename]")
		return
	}

	command := os.Args[1]
	filename := ""
	if len(os.Args) > 2 {
		filename = os.Args[2]
	}

	switch command {
	case "push":
		content := readContent(filename)
		if content == "" {
			fmt.Println("No content to push.")
			return
		}
		err := postContent(config.serverUrl, content)
		if err != nil {
			fmt.Printf("Failed to push content: %v\n", err)
			return
		}
		fmt.Println("Content pushed successfully.")
	case "pull":
		content, err := getContent(config.serverUrl)
		if err != nil {
			fmt.Printf("Failed to pull content: %v\n", err)
			return
		}
		if filename == "" {
			fmt.Println(content)
		} else {
			err := saveContent(content, filename)
			if err != nil {
				fmt.Printf("Failed to save content: %v\n", err)
				return
			}
			fmt.Println("Content saved successfully.")
		}
	default:
		fmt.Println("Invalid command.")
	}
}

func readContent(filename string) string {
	if filename == "" {
		// Read from stdin
		content, _ := io.ReadAll(os.Stdin)
		return string(content)
	}

	// Read from file
	content, _ := os.ReadFile(filename)
	return string(content)
}

func postContent(baseUrl, content string) error {
	url := baseUrl + "/api/v1/update"
	payload := fmt.Sprintf(`{"content": "%s"}`, content)
	_, err := http.Post(url, "application/json", bytes.NewBufferString(payload))
	return err
}

func getContent(baseUrl string) (string, error) {
	url := baseUrl + "/api/v1/content"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	content, _ := io.ReadAll(resp.Body)
	return string(content), nil
}

func saveContent(content string, filename string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}

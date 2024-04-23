package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"todo/config"
	"todo/lib"
)

func main() {
	config := lib.LoadConfig[config.Config]()
	if len(os.Args) < 2 {
		log.Println("Usage: client [push|pull] [filename]")
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
			log.Println("No content to push.")
			return
		}
		err := postContent(config.ServerUrl, config.Token, content)
		if err != nil {
			log.Printf("Failed to push content: %v\n", err)
			return
		}
		log.Println("Content pushed successfully.")
	case "pull":
		content, err := getContent(config.ServerUrl, config.Token)
		if err != nil {
			log.Printf("Failed to pull content: %v\n", err)
			return
		}
		if filename == "" {
			log.Println(content)
		} else {
			err := saveContent(content, filename)
			if err != nil {
				log.Printf("Failed to save content: %v\n", err)
				return
			}
			log.Println("Content saved successfully.")
		}
	default:
		log.Println("Invalid command.")
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

func postContent(baseUrl, token, content string) error {
	payload, _ := json.Marshal(map[string]string{"content": content})
	req, err := http.NewRequest("POST", baseUrl+"/api/v1/update", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func getContent(baseUrl, token string) (string, error) {
	req, err := http.NewRequest("GET", baseUrl+"/api/v1/content", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var data map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	return data["content"], nil
}

func saveContent(content string, filename string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}

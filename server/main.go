package main

import (
	"log"
	"net/http"
	"os"
	"todo/lib"

	"github.com/gin-gonic/gin"
)

type Config struct {
	lib.ServerConfig
}

func main() {
	config := lib.LoadConfig[Config]()

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/update", func(c *gin.Context) {
		var request struct {
			Content string `json:"content"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		saveToServer(request.Content)
		c.JSON(http.StatusOK, gin.H{"content": request.Content})
	})
	api.GET("/content", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"content": getContent()})
	})

	log.Fatal(router.Run(config.ServerConfig.Port))
}

func getContent() string {
	content, err := os.ReadFile("data.md")
	if err == nil {
		return string(content)
	}
	if os.IsNotExist(err) {
		err = os.WriteFile("data.md", []byte(""), 0644)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
	return ""
}

func saveToServer(content string) {
	err := os.WriteFile("data.md", []byte(content), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

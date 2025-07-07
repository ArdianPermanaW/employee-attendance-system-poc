package main

import (
	"gin/internal/config"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	dbConn, err := config.Connect(cfg.DB)
	print(dbConn)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run() // listen and serve on 0.0.0.0:8080
}

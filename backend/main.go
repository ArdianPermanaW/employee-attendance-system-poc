package main

import (
	"gin/internal/config"
	"gin/internal/employees"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}
	//connect to DB
	db, err := config.Connect(cfg.DB)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	// Auto-migrate
	db.AutoMigrate(&employees.Employee{})

	// Init layer
	repo := employees.NewRepository(db)
	service := employees.NewService(repo)
	handler := employees.NewHandler(service)

	// Set up routes
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	employees.RegisterRoutes(router, handler)

	router.Run(":" + cfg.Port)
}

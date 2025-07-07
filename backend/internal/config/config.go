package config

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Logs LogConfig
	DB   PostgresConfig
	Port string
}

type LogConfig struct {
	Style string
	Level string
}

type PostgresConfig struct {
	Username string
	Password string
	URL      string
	Port     string
	Name     string
}

func LoadConfig() (*Config, error) {

	cfg := &Config{
		Port: os.Getenv("PORT"),
		Logs: LogConfig{
			Style: os.Getenv("LOG_STYLE"),
			Level: os.Getenv("LOG_LEVEL"),
		},
		DB: PostgresConfig{
			Username: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			URL:      os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Name:     os.Getenv("DB_NAME"),
		},
	}
	fmt.Println(cfg.Port, cfg.DB.URL, cfg.DB.Username, cfg.DB.Password, cfg.DB.Name, cfg.DB.Port, "oob")
	return cfg, nil
}

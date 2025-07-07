package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg PostgresConfig) (*gorm.DB, error) {
	fmt.Println(cfg.URL, cfg.Username, cfg.Password, cfg.Name, cfg.Port, "boo")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.URL, cfg.Username, cfg.Password, cfg.Name, cfg.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}

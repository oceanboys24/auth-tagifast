package configs

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/oceanboys24/auth-tagifast/internal/auth/types"
)

func LoadEnv() (*types.ConfigDB, error) {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	cfg := &types.ConfigDB{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	log.Println("Config loaded successfully")
	return cfg, nil
}

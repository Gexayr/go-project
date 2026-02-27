package configs

import (
	"log"
	"os"

	"github.com/lpernett/godotenv"
)

type Config struct {
	Db   DbConfig
	Auth AuthConfig
}

type DbConfig struct {
	Dsn string
}

type AuthConfig struct {
	Secret string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		Db: DbConfig{
			os.Getenv("DSN"),
		},
		Auth: AuthConfig{
			os.Getenv("SECRET"),
		},
	}
}

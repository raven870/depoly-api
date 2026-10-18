package config

import (
	"log"
	"log/slog"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	DB_URI string `env:"DB_URI"`
	Env    string `env:"ENV"`
	PORT   string `env:"PORT"`
}

func ConfigureEnv() *Config {

	error := godotenv.Load()
	if error != nil {
		slog.Info("Error loading .env file", "message:", error.Error())
	}

	var config Config
	if err := env.Parse(&config); err != nil {
		log.Fatal("Error While Parsing .env file", err.Error())
	}

	return &config

}

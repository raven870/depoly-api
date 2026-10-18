package config

import (
	"log"

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
		log.Fatal("Error While loading .env file", error.Error())
	}

	var config Config
	if err := env.Parse(&config); err != nil {
		log.Fatal("Error While Parsing .env file", err.Error())
	}

	return &config

}

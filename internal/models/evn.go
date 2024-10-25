package models

import (
	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
	"log"
)

const filePath = ".env.local"

type Environment struct {
	RconnAddress  string `env:"RCONN_ADRESS,required,notEmpty"`
	RconnPassword string `env:"RCONN_PASSWORD,required,notEmpty"`
}

func ParseEvn() *Environment {
	envAdres := Environment{}

	if err := godotenv.Load(filePath); err != nil {
		log.Fatal("Error loading .env file")
	}
	if err := env.Parse(&envAdres); err != nil {
		log.Fatal(err)

	}
	return &envAdres
}

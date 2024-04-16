package Database

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadENVVar() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}
}

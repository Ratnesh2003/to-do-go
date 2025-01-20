package configs

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVars() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Some error occured while loading environment variables.")
    }
}
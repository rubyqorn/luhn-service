package utils

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func GetEnvVariable(variableName string) string {
	path, err := filepath.Abs("../.env")

	if err != nil {
		log.Fatalf("Env file not found")
	}

	loadErr := godotenv.Load(path)

	if loadErr != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(variableName)
}

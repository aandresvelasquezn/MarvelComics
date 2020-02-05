package tools

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//GetDotEnvVariable Get value from Key in .env
func GetDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load("configs/.envdev")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

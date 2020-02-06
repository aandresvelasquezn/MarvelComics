package tools

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//DotEnvVariable Get value from Key in .env
func DotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load("configs/.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

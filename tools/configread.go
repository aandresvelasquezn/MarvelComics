package tools

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//DotEnvVariable Get value from Key in .env
func DotEnvVariable(key string) string {

	// load .env file
	_, err := os.Stat("configs/.env")
	if !os.IsNotExist(err) {
		err = godotenv.Load(os.ExpandEnv("configs/.env"))
	}

	if err != nil {
		log.Println(err)
	}

	return os.Getenv(key)
}

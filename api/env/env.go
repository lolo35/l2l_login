package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Env(env_value string) string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	return os.Getenv(env_value)
}

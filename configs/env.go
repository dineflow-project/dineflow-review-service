package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadDotEnv() {
	env := os.Getenv("USE_DOT_ENV")
	// if flag not set load .env file
	if env == "" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

func EnvMongoURI() string {
	loadDotEnv()
	return os.Getenv("MONGO_URI")
}

func EnvMongoDBName() string {
	loadDotEnv()
	return os.Getenv("MONGO_DATABASE_NAME")
}

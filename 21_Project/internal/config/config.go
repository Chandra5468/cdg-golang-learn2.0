package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
}

func (c *Config) Conf() {
	env := os.Getenv("APP_ENV")

	if env == "" {
		env = "local"
	}

	envFile := fmt.Sprintf("config/.env.%s", env)

	err := godotenv.Load(envFile)

	if err != nil {
		log.Fatalf("error loading %s file: %s", envFile, err.Error())
	}

	// PORT := os.Getenv("port")

}

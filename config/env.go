package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	PublicHost string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

var EnvVars = initConfig()

func initConfig() *Config {
	return &Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "mypassword"),
		DBAddress: fmt.Sprintf("%s:%s",
			getEnv("DB_HOST", "localhost"), getEnv("DB_PORT", "3306")),
		DBName: getEnv("DB_NAME", "ecom"),
	}
}

func getEnv(key, fallback string) string {
	_ = godotenv.Load()

	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return val
}

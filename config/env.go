package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	PublicHost             string
	DBUser                 string
	DBPassword             string
	DBAddress              string
	DBName                 string
	JWTExpirationInSeconds int64
	JWTSecret              string
}

var EnvVars = initConfig()

func initConfig() *Config {
	return &Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "mypassword"),
		DBAddress: fmt.Sprintf("%s:%s",
			getEnv("DB_HOST", "localhost"), getEnv("DB_PORT", "3306")),
		DBName:                 getEnv("DB_NAME", "ecom"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXP", 3600*24),
		JWTSecret:              getEnv("JWT_SECRET", "86400secondstohours86400secondstohours"),
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

func getEnvAsInt(key string, fallback int64) int64 {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	valAsInt, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return fallback
	}

	return valAsInt
}

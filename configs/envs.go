package configs

import (
	"fmt"
	"os"
)

type Config struct {
	PublictHost string
	Port        string
	DBUser      string
	DBPassword  string
	DBAddress   string
	DBName      string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		PublictHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:        getEnv("PORT", "8080"),
		DBUser:      getEnv("DB_USER", "root"),
		DBPassword:  getEnv("DB_PASSWORD", "root"),
		DBAddress:   fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "8889")),
		DBName:      getEnv("DB_NAME", "portfolio_cms"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

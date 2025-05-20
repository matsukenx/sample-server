package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Port    string
	APIKeys map[string]string
}

func LoadFromEnv() *Config {
	_ = godotenv.Load()

	port := os.Getenv("HTTP_SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	apiKeys := make(map[string]string)
	for _, pair := range strings.Split(os.Getenv("API_KEYS"), ",") {
		kv := strings.SplitN(pair, ":", 2)
		if len(kv) == 2 {
			apiKeys[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
		}
	}

	return &Config{
		Port:    port,
		APIKeys: apiKeys,
	}
}

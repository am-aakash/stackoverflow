package config

import (
	"fmt"
	"os"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
	ServerPort string
}

func LoadConfig() *Config {
	return &Config{
			DBHost:     getEnvOrDefault("DB_HOST", "localhost"),
			DBPort:     getEnvOrDefault("DB_PORT", "3306"),
			DBUser:     getEnvOrDefault("DB_USER", "root"),
			DBPassword: getEnvOrDefault("DB_PASSWORD", "root"),
			DBName:     getEnvOrDefault("DB_NAME", "stackoverflow_clone"),
			JWTSecret:  getEnvOrDefault("JWT_SECRET", "your-secret-key"),
			ServerPort: getEnvOrDefault("SERVER_PORT", "8080"),
	}
}

func (c *Config) GetDBConnString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True",
			c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
			return value
	}
	return defaultValue
}
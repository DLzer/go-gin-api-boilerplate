package configs

import (
	"fmt"
	"os"
)

// DBConfig object
type DBConfig struct {
	Username string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Hostname string `env:"DB_HOST"`
	Database string `env:"DB_NAME"`
	Port     string `env:"DB_PORT"`
	DSN      string
}

// GetDBConfig returns the populated DBConfig object
func GetDBConfig() DBConfig {

	var username = os.Getenv("DB_USER")
	var password = os.Getenv("DB_PASSWORD")
	var hostname = os.Getenv("DB_HOST")
	var database = os.Getenv("DB_NAME")
	var port = os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		hostname, port, username, password, database)

	return DBConfig{
		Username: username,
		Password: password,
		Hostname: hostname,
		Database: database,
		Port:     port,
		DSN:      dsn,
	}
}

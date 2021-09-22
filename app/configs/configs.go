package configs

import "os"

const (
	prod = "production"
)

// Config object
type Config struct {
	Env   string   `env:"APP_ENV"`
	DB    DBConfig `json:"db"`
	Host  string   `env:"APP_HOST"`
	Port  string   `env:"APP_PORT"`
	Token string   `env:"APP_TOKEN"`
}

// IsProd Checks if env is production
func (c Config) IsProd() bool {
	return c.Env == prod
}

// GetConfig gets all config for the application
func GetConfig() Config {
	return Config{
		Env:  os.Getenv("APP_ENV"),
		DB:   GetDBConfig(),
		Host: os.Getenv("APP_HOST"),
		Port: os.Getenv("APP_PORT"),
		// Token:   os.Getenv("APP_TOKEN"),
	}
}

package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment string // develop, staging, production

	PostgresHost           string
	PostgresPort           int
	PostgresUser           string
	PostgresPassword       string
	PostgresDB             string
	PostgresMigrationsPath string
	// context timeout in seconds
	CtxTimeout int

	StoreHost string

	LogLevel string
	HTTPPort string

	TelegramApiToken string
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDB = cast.ToString(getOrReturnDefault("POSTGRES_DB", ""))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", ""))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", ""))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":9000"))
	c.StoreHost = cast.ToString(getOrReturnDefault("TASK_SERVICE_HOST", "localhost"))

	c.CtxTimeout = cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7))
	c.TelegramApiToken = cast.ToString(getOrReturnDefault("TELEGRAM_API_TOKEN", ""))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}

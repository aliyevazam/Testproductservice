package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...

type Config struct {
	Environment       string //develop, staging, production
	PostgresHost      string
	PostgresPort      int
	PostgresDatabase  string
	PostgresUser      string
	PostgresPassword  string
	LogLevel          string
	RPCPort           string
	ReviewServiceHost string
	ReviewServicePort int
	StoreServiceHost string
	StoreServicePort int
}

// Load loads environment vars and inflates Config

func Load() Config {
	c := Config{}
	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "testdb"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "azam"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "Azam_2000"))
	
	c.StoreServiceHost = cast.ToString(getOrReturnDefault("Stores_host","localhost"))
	c.StoreServicePort = cast.ToInt(getOrReturnDefault("Stores_port",9000))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.RPCPort = cast.ToString(getOrReturnDefault("RPC_PORT", ":8000"))
	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}

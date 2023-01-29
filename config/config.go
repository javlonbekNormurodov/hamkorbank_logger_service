package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	// DebugMode indicates service mode is debug.
	DebugMode = "debug"
	// TestMode indicates service mode is test.
	TestMode = "test"
	// ReleaseMode indicates service mode is release.
	ReleaseMode = "release"
	// listening only errors
	AllErrors = "consumer1.error"
	// listening only info
	AllInfo = "consumer2.info"
	// listening only debug
	AllDebug = "consumer3.debug"
	All      = "#" // listening all
)

type Config struct {
	ServiceName string
	Environment string // debug, test, release
	Version     string

	ServiceHost string
	HTTPPort    string
	HTTPScheme  string

	DefaultOffset string
	DefaultLimit  string

	RabbitMQHost     string
	RabbitMQPort     int
	RabbitMQUser     string
	RabbitMQPassword string

	LogLevel string

	PostgresMaxConnections int32
}

// Load ...
func Load() Config {
	if err := godotenv.Load("/app/.env"); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.LogLevel = cast.ToString(getOrReturnDefaultValue("LOG_LEVEL", "debug"))

	config.ServiceName = cast.ToString(getOrReturnDefaultValue("SERVICE_NAME", "epa_go_api_gateway"))
	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", DebugMode))
	config.Version = cast.ToString(getOrReturnDefaultValue("VERSION", "1.0"))

	config.ServiceHost = cast.ToString(getOrReturnDefaultValue("SERVICE_HOST", "localhost"))
	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":8000"))
	config.HTTPScheme = cast.ToString(getOrReturnDefaultValue("HTTP_SCHEME", "http"))

	config.PostgresMaxConnections = cast.ToInt32(getOrReturnDefaultValue("POSTGRES_MAX_CONNECTIONS", 30))

	config.DefaultOffset = cast.ToString(getOrReturnDefaultValue("DEFAULT_OFFSET", "0"))
	config.DefaultLimit = cast.ToString(getOrReturnDefaultValue("DEFAULT_LIMIT", "10000000"))

	config.RabbitMQUser = cast.ToString(getOrReturnDefaultValue("RABBIT_MQ_USER", "guest"))
	config.RabbitMQPassword = cast.ToString(getOrReturnDefaultValue("RABBIT_MQ_PASSWORD", "guest"))
	config.RabbitMQHost = cast.ToString(getOrReturnDefaultValue("RABBIT_MQ_HOST", "localhost"))
	config.RabbitMQPort = cast.ToInt(getOrReturnDefaultValue("RABBIT_MQ_PORT", 5672))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}

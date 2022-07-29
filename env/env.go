package env

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnv() map[string]string {
	listOfEnv := []string{
		"PROCESS_NAME",
		"API_PORT", "API_PATH",
		"RPC_PORT",
		"DB_HOST", "DB_PORT", "DB_USER", "DB_PASS", "DB_NAME", "DB_POOL",
		"LOG_LEVEL",
		"FILE_LOG",
		"KAFKA_LOG", "KAFKA_BROKER", "KAFKA_TOPIC",
	}

	env, _ := godotenv.Read(".env")

	for _, envVar := range listOfEnv {
		if val, ok := os.LookupEnv(envVar); ok {
			env[envVar] = val
		}
	}
	return env
}

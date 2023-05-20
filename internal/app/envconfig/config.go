package envConfig

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	DBConfig DBConfig
}

type DBConfig struct {
	DBPwd string
}

func New() *Config {
	return &Config{
		DBConfig: DBConfig{
			DBPwd: getEnv("DB_PWD", ""),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(key string, defaultVal int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	logrus.Info(fmt.Sprintf("cant get env variable %s as int, returned default...", key))
	return defaultVal
}

func getEnvAsBool(key string, defaultVal bool) bool {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}
	logrus.Info(fmt.Sprintf("cant get env variable %s as bool, returned default...", key))
	return defaultVal
}

func getEnvAsSlice(key string, defaultVal []string, sep string) []string {
	valStr := getEnv(key, "")

	if valStr == "" {
		logrus.Info(fmt.Sprintf("cant get env variable %s as slice, returned default...", key))
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}

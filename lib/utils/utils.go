package utils

import (
	"os"
	"strconv"

	"github.com/spf13/viper"
)

func GetString(key string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}

	return viper.GetString(key)
}

func GetInt(key string) int {
	if value, exist := os.LookupEnv(key); exist {
		valInt, _ := strconv.Atoi(value)
		return valInt
	}

	return viper.GetInt(key)
}

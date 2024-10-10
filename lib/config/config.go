package config

import (
	"file-server/lib/utils"

	"github.com/spf13/viper"
)

type Config struct {
	APIConfig APIConfig
	APPConfig APPConfig
}

type APPConfig struct {
	HTTPPort int
	BaseURL  string
}

func getAPPConfig() APPConfig {
	return APPConfig{
		HTTPPort: utils.GetInt("HTTP_PORT"),
		BaseURL:  utils.GetString("BASE_URL"),
	}
}

type APIConfig struct {
	AuthURL string
}

func getAPIConfig() APIConfig {
	return APIConfig{
		AuthURL: utils.GetString("AUTH_URL"),
	}
}

func GetConfig(path, filename, fileType string) *Config {
	viper.SetConfigName(filename)
	viper.SetConfigType(fileType)
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	return &Config{
		APIConfig: getAPIConfig(),
		APPConfig: getAPPConfig(),
	}
}

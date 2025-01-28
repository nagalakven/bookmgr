package config

import (
	"bookmgr/pkg/logger"
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Application struct {
		Name string `mapstructure:"name"`
	} `mapstructure:"application"`

	HTTP struct {
		Port           int `mapstructure:"port"`
		ReadTimeout    int `mapstructure:"readTimeout"`
		WriteTimeout   int `mapstructure:"writeTimeout"`
		MaxHeaderBytes int `mapstructure:"maxHeaderBytes"`
	} `mapstructure:"http"`
}

type ConfigInterface interface {
	LoadConfig(path string)
}

var AppConfig Config

// Load configuration from yaml
func LoadConfig(path string) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		logger.LogMessage(logger.FATAL, fmt.Sprintf("Error reading config file, %s", err), false)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		logger.LogMessage(logger.FATAL, fmt.Sprintf("Unable to decode into struct, %v", err), false)
	}
}

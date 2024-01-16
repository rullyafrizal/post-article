package config

import (
	"github.com/spf13/viper"
)

var config *Config

func Init() error {
	// Config File Path
	viper.AddConfigPath("./config/")
	// Config File Name
	viper.SetConfigName("config")
	// Config File Type
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	config = new(Config)
	err = viper.Unmarshal(&config)
	if err != nil {
		return err
	}

	return nil
}

func Get() *Config {
	if config == nil {
		config = &Config{}
	}
	return config
}

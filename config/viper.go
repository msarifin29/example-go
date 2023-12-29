package config

import "github.com/spf13/viper"

func NewViper() *viper.Viper {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath("..")

	return config
}

package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Name        string
	Environment string
}

func Newconfig() *Config {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	instance := new(Config)
	err := viper.Unmarshal(&instance)
	if err != nil {
		fmt.Println(err)
	}

	return instance

}

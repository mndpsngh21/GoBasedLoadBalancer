package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() map[string]interface{} {
	fmt.Println("loading configuration")
	viper.SetConfigName("servers")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicf("Fatal error config file: %s\n", err)
	}
	return viper.AllSettings()
}

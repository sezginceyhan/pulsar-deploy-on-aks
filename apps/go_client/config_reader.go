package main

import (
	"github.com/spf13/viper"
)

var C Config

type Config struct {
	Pulsar PulsarConfig `yaml:"Pulsar"`
}

type PulsarConfig struct {
	Url  string `yaml:"url"`
	Port string `yaml:"port"`
}

func ReadConfigFile() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("Config not found")
		} else {
			log.Fatal("Err", err)
		}
	}

	err := viper.Unmarshal(&C)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

}

package config

import (
	"errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func NewConfig() (*Config, error) {
	var found bool
	path := []string{
		"/etc/api-config",
		"files/etc/api-config",
		"./files/etc/api-config",
		"../files/etc/api-config",
		"../../files/etc/api-config",
	}

	viper.SetConfigName("cfg.example")
	viper.SetConfigType("yaml")
	for _, v := range path {
		viper.AddConfigPath(v)
		err := viper.ReadInConfig() // Find and read the config file
		if err != nil {             // Handle errors reading the config file
			log.Error().Err(err)
		} else {
			found = true
			break
		}
	}

	if !found {
		return nil, errors.New("no config file found")
	}

	cfg := &Config{}
	err := viper.Unmarshal(cfg)
	return cfg, err
}

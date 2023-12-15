package config

import (
	"github.com/spf13/viper"
)

const ConfigName = "config"
const ConfigType = "yaml"

var Configuration *Config

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
}

func GetConfig() Config {
	if Configuration == nil {
		err := LoadConfig("./6singleton")
		if err != nil {
			panic(err)
		}
	}
	return *Configuration
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(ConfigName)
	viper.SetConfigType(ConfigType)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	var config Config
	err = viper.Unmarshal(&config)
	Configuration = &config
	return
}

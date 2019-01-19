package config

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/erraa/doninja/utils"
)

var log = utils.LogWithPrefix("config")
var config Config

type Config struct {
	Discord Discord
}

type Discord struct {
	Token     string
	BotPrefix string
}

// ReadConfig read
func ReadConfig() Config {
	return config
}

func init() {
	log.Info("Reading new file...")

	viper.SetConfigName("doninja")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/")
	viper.AddConfigPath("$HOME/")

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Error reading config", err))
	}

	err = viper.Unmarshal(&config)

	if err != nil {
		panic(fmt.Errorf("Error reading config", err))
	}

}

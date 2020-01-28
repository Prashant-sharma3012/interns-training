package helpers

import (
	"log"

	"github.com/interns-training/models"
	"github.com/spf13/viper"
)

func LoadConfig() *models.Config {
	//load config
	config := &models.Config{}

	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Panicf("Fatal error config file: %s \n", err)
	}

	viper.Unmarshal(config)

	return config
}

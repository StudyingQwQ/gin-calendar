package config

import (
	"github.com/spf13/viper"
)

var Config *viper.Viper

func InitConfig() {
	viper.SetConfigFile("etc/config.yaml")
	viper.ReadInConfig()
	Config = viper.GetViper()
}

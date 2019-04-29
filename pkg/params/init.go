package params

import (
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("listeningport", "8000")
	viper.SetDefault("rooturl", "/")
	viper.SetDefault("delay", "300")
	viper.SetDefault("data_dir", "")
	viper.SetDefault("logging", "json")
	viper.SetDefault("nodeid", "0")
	viper.BindEnv("listeningport")
	viper.BindEnv("rooturl")
	viper.BindEnv("logging")
	viper.BindEnv("nodeid")
}
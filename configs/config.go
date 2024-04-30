package configs

import (
	"log"

	"github.com/spf13/viper"
)

var conf *Config

func Init(path string) {
	if path == "" {
		path = "configs"
	}
	initViper(path)
	loadConfigs()
}

func GetConfig() Config {
	if conf == nil {
		loadConfigs()
	}
	return *conf
}

func initViper(path string) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("cannot read config file: %s", err)
	}

}

func loadConfigs() {
	if err := viper.Unmarshal(&conf); err != nil {
		log.Panicf("load config: %v", err)
		return
	}

	if err := conf.Validate(); err != nil {
		log.Panicf("validate config: \n%v\n", err)
	}
}

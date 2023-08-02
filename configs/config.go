package configs

import (
	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	DBDrive string
	DBHost  string
	DBName  string
	DBPort  string
	DBUser  string
	DBPass  string

	ServerPort string

	JWTSecret    string
	JWTExpiresIn int
}

func Init(path string) (*config, error) {
	viper.SetConfigFile("config_app")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, err
}

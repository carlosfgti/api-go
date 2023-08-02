package configs

import (
	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	DBDrive string `mapstructure:"DB_DRIVE"`
	DBHost  string `mapstructure:"DB_HOST"`
	DBName  string `mapstructure:"DB_NAME"`
	DBPort  string `mapstructure:"DB_PORT"`
	DBUser  string `mapstructure:"DB_USER"`
	DBPass  string `mapstructure:"DB_PASS"`

	ServerPort string `mapstructure:"SERVER_PORT"`

	JWTSecret    string `mapstructure:"JWT_SECRET"`
	JWTExpiresIn int    `mapstructure:"JWT_EXPIRES_IN"`
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

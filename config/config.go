package config

import "github.com/spf13/viper"

// Config struct maps to the environment variables
type Config struct {
	SECRETKEY   string `mapstructure:"JWTSECRET"`
	DBHost        string `mapstructure:"DBHOST"`
	DBUser        string `mapstructure:"DBUSER"`
	DBPassword    string `mapstructure:"DBPASSWORD"`
	DBDatabase    string `mapstructure:"DBNAME"`
	DBPort        string `mapstructure:"DBPORT"`
	DBSslmode     string `mapstructure:"DBSSL"`
}

// LoadConfig will load the environment variables to accessible.
func LoadConfig() *Config {
	var cnfg Config
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.Unmarshal(&cnfg)
	return &cnfg
}
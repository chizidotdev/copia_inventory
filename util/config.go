package util

import (
	"log"

	"github.com/spf13/viper"
)

// Config stores all the configuration for the application
// using values read from the config file or env variables
type Config struct {
	DBDriver           string `mapstructure:"DB_DRIVER"`
	DBSource           string `mapstructure:"DB_SOURCE"`
	PORT               string `mapstructure:"PORT"`
	AuthSecret         string `mapstructure:"AUTH_SECRET"`
	AuthDomain         string `mapstructure:"AUTH_DOMAIN"`
	AuthCallbackURL    string `mapstructure:"AUTH_CALLBACK_URL"`
	GoogleClientID     string `mapstructure:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret string `mapstructure:"GOOGLE_CLIENT_SECRET"`
}

var EnvVars Config

// LoadConfig loads the configuration from the config file or env variable
func LoadConfig() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Cannot read config file:", err)
	}
	viper.AutomaticEnv()

	EnvVars.DBDriver = viper.GetString("DB_DRIVER")
	EnvVars.DBSource = viper.GetString("DB_SOURCE")
	EnvVars.PORT = viper.GetString("PORT")

	EnvVars.AuthSecret = viper.GetString("AUTH_SECRET")
	EnvVars.AuthDomain = viper.GetString("AUTH_DOMAIN")
	EnvVars.AuthCallbackURL = viper.GetString("AUTH_CALLBACK_URL")
	EnvVars.GoogleClientID = viper.GetString("GOOGLE_CLIENT_ID")
	EnvVars.GoogleClientSecret = viper.GetString("GOOGLE_CLIENT_SECRET")
}

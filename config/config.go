package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	IapHost                   string `mapstructure:"iap_host"`
	ServiceAccountCredentials string `mapstructure:"service_account_credentials"`
	ClientID                  string `mapstructure:"client_id"`
	LoggerLevel               string `mapstructure:"logger_level"`
	SentryDSN                 string `mapstructure:"sentry_dsn"`
	RefreshTimeSeconds        string `mapstructure:"refresh_time_seconds"`
	Port                      string `mapstructure:"port"`
}

func Load() (Config, error) {
	viper.SetDefault("LOGGER_LEVEL", "error")
	viper.BindEnv("iap_host")
	viper.BindEnv("service_account_credentials")
	viper.BindEnv("client_id")
	viper.BindEnv("logger_level")
	viper.BindEnv("sentry_dsn")
	viper.BindEnv("refresh_time_seconds")
	viper.BindEnv("port")

	viper.SetConfigName("iap.conf")

	viper.AddConfigPath("/etc/iap_auth/")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	viper.ReadInConfig()
	viper.AutomaticEnv()

	var c Config
	err := viper.Unmarshal(&c)
	if err != nil {
		return Config{}, err
	}
	return c, nil
}

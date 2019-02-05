package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	IapHost                   string        `mapstructure:"iap_host"`
	ServiceAccountCredentials string        `mapstructure:"service_account_credentials"`
	ClientID                  string        `mapstructure:"client_id"`
	LoggerLevel               string        `mapstructure:"logger_level"`
	SentryDSN                 string        `mapstructure:"sentry_dsn"`
	RefreshTime               time.Duration `mapstructure:"refresh_time"`
}

func Load() (Config, error) {
	viper.SetDefault("LOGGER_LEVEL", "error")

	viper.SetConfigName("iap.conf")

	viper.AddConfigPath("/etc/iap_auth/")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")

	viper.AutomaticEnv()
	viper.ReadInConfig()
	var c Config
	err := viper.Unmarshal(&c)
	if err != nil {
		return Config{}, err
	}
	return c, nil
}

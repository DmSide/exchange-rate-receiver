package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct {
	GRPCPort    string `mapstructure:"GRPC_PORT"`
	DatabaseURL string `mapstructure:"DATABASE_URL"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // required if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	viper.AutomaticEnv()          // read in environment variables that match

	viper.SetDefault("GRPCPort", "50051") // default value for GRPC port

	if err := viper.ReadInConfig(); err != nil {
		// TODO: Add special name here
		logger, _ := zap.NewProduction()
		logger.Warn("No config file found", zap.Error(err))
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

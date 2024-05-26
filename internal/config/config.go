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
	logger, _ := zap.NewProduction()

	viper.BindEnv("GRPC_PORT")
	viper.BindEnv("DATABASE_URL")

	viper.AutomaticEnv()

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	// Please remove it for PROD
	logger.Info("Configuration loaded", zap.String("GRPC_PORT", cfg.GRPCPort), zap.String("DATABASE_URL", cfg.DatabaseURL))

	return &cfg, nil
}

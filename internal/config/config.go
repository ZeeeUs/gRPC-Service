package config

import (
	"time"

	"github.com/spf13/viper"
)

type GRPCConfig struct {
	Network           string
	Address           string
	MaxConnectionIdle time.Duration
	Timeout           time.Duration
	MaxConnectionAge  time.Duration
}

type Config struct {
	GRPCConfig GRPCConfig
}

func SetupSource(cfgType string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigType(cfgType)
	v.AutomaticEnv()

	return v, nil
}

func GetConfig() (*Config, error) {
	v, err := SetupSource("env")
	if err != nil {
		return nil, err
	}

	return &Config{
		GRPCConfig: GRPCConfig{
			Network:           v.GetString("GRPC_NETWORK"),
			Address:           v.GetString("GRPC_ADDRESS"),
			MaxConnectionIdle: v.GetDuration("GRPC_MAX_CONN_IDLE"),
			Timeout:           v.GetDuration("GRPC_SERVER_TIMEOUT"),
			MaxConnectionAge:  v.GetDuration("GRPC_MAX_CONN_AGE"),
		},
	}, nil
}

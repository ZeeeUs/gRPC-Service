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

type DatabaseConfig struct {
	Name            string
	Address         string
	User            string
	Password        string
	MaxIdleLifetime time.Duration
	MaxLifetime     time.Duration
	PrepareCacheCap int
	MaxConn         int
}

type Config struct {
	GRPCConfig     GRPCConfig
	DatabaseConfig DatabaseConfig
}

func SetupSource(cfgType string) *viper.Viper {
	v := viper.New()

	v.SetDefault("PG_MAX_IDLE_LIFETIME", "30s")
	v.SetDefault("PG_MAX_LIFETIME", "3m")
	v.SetDefault("PG_PREPARE_CACHE_CAPACITY", 128)
	v.SetDefault("PG_MAX_CONNECTIONS", 15)

	v.SetConfigType(cfgType)
	v.AutomaticEnv()

	return v
}

func GetConfig() *Config {
	v := SetupSource("env")

	return &Config{
		GRPCConfig: GRPCConfig{
			Network:           v.GetString("GRPC_NETWORK"),
			Address:           v.GetString("GRPC_ADDRESS"),
			MaxConnectionIdle: v.GetDuration("GRPC_MAX_CONN_IDLE"),
			Timeout:           v.GetDuration("GRPC_SERVER_TIMEOUT"),
			MaxConnectionAge:  v.GetDuration("GRPC_MAX_CONN_AGE"),
		},
		DatabaseConfig: DatabaseConfig{
			Name:            v.GetString("PG_DATABASE_NAME"),
			Address:         v.GetString("PG_ADDRESS"),
			User:            v.GetString("PG_USER"),
			Password:        v.GetString("PG_PASSWORD"),
			MaxIdleLifetime: v.GetDuration("PG_MAX_IDLE_LIFETIME"),
			MaxLifetime:     v.GetDuration("PG_MAX_LIFETIME"),
			PrepareCacheCap: v.GetInt("PG_PREPARE_CACHE_CAPACITY"),
			MaxConn:         v.GetInt("PG_MAX_CONNECTIONS"),
		},
	}
}

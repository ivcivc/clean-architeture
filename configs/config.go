package configs

import (
	"github.com/spf13/viper"
)

type conf struct {
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBHost            string `mapstructure:"DB_HOST"`
	DBPort            string `mapstructure:"DB_PORT"`
	DBUser            string `mapstructure:"DB_USER"`
	DBPassword        string `mapstructure:"DB_PASSWORD"`
	DBName            string `mapstructure:"DB_NAME"`
	WebServerPort     string `mapstructure:"WEB_SERVER_PORT"`
	GRPCServerPort    string `mapstructure:"GRPC_SERVER_PORT"`
	GraphQLServerPort string `mapstructure:"GRAPHQL_SERVER_PORT"`
}

func LoadConfig(path string) (*conf, error) {
	var cfg *conf

	// Set default values
	viper.SetDefault("DB_DRIVER", "mysql")
	viper.SetDefault("DB_HOST", "mysql")
	viper.SetDefault("DB_PORT", "3308")
	viper.SetDefault("DB_USER", "root")
	viper.SetDefault("DB_PASSWORD", "root")
	viper.SetDefault("DB_NAME", "orders")
	viper.SetDefault("WEB_SERVER_PORT", ":8080")
	viper.SetDefault("GRPC_SERVER_PORT", "50051")
	viper.SetDefault("GRAPHQL_SERVER_PORT", "8081")

	viper.AutomaticEnv()

	// Try to read .env file if it exists, but don't panic if it doesn't
	viper.SetConfigFile(".env")
	viper.ReadInConfig() // Ignore any errors - use env vars or defaults

	err := viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, nil
}

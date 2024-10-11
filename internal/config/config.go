package config

import (
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var config *viper.Viper

func init() {
	config = viper.New()
	config.SetConfigName("grpcexample_config")
	config.SetConfigType("yaml")
	config.AddConfigPath("/run/secrets")

	err := config.ReadInConfig()
	if err != nil {
		fmt.Println("Failed to load configuration")
	}

	config.SetDefault("port", "8080")
	config.SetDefault("host", "")
}

func GetHost() string {
	const host = "host"

	return config.GetString(host)
}

func GetPort() string {
	const port = "port"

	return config.GetString(port)
}

func GetZapConfig() (zap.Config, error) {
	const zapConfig = "zap-config"

	if !config.IsSet(zapConfig) {
		fmt.Println("zap-config is not specified. The production mode is automatically selected")
		return zap.NewProductionConfig(), nil
	}
	config.GetString(zapConfig)
	switch config.GetString(zapConfig) {
	case "prod":
		return zap.NewProductionConfig(), nil
	case "dev":
		return zap.NewDevelopmentConfig(), nil
	default:
		fmt.Println("zap-config is specified incorrectly. The production mode is automatically selected")
		return zap.NewProductionConfig(), nil
	}
}

func GetServer() string {
	host := GetHost()

	port := GetPort()

	return host + ":" + port
}

func GetPsqlConnection() (string, error) {
	const (
		postgres      = "postgres"
		hostField     = "host"
		portField     = "port"
		userField     = "user"
		passwordField = "password"
		dbNameField   = "db-name"
	)

	postgresConfig := config.Sub(postgres)

	if !postgresConfig.IsSet(hostField) {
		zap.S().Errorw("Error while getting host from postgres config", ErrNoKey)
		return "", ErrNoKey
	}
	host := postgresConfig.GetString(hostField)

	if !postgresConfig.IsSet(portField) {
		zap.S().Errorw("Error while getting port from postgres config", ErrNoKey)
		return "", ErrNoKey
	}
	port := postgresConfig.GetInt(portField)

	if !postgresConfig.IsSet(userField) {
		zap.S().Errorw("Error while getting user from postgres config", ErrNoKey)
		return "", ErrNoKey
	}
	user := postgresConfig.GetString(userField)

	if !postgresConfig.IsSet(passwordField) {
		zap.S().Errorw("Error while getting password from postgres config", ErrNoKey)
		return "", ErrNoKey
	}
	password := postgresConfig.GetString(passwordField)

	if !postgresConfig.IsSet(dbNameField) {
		zap.S().Errorw("Error while getting database name from postgres config", ErrNoKey)
		return "", ErrNoKey
	}
	dbname := postgresConfig.GetString(dbNameField)

	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		host, port, user, password, dbname), nil
}

package server

import "gophkeeper/pkg/logger"

type Config struct {
	Logger logger.Config `mapstructure:"log"`
}

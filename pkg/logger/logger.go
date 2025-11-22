package logger

import "go.uber.org/zap"

var logger *zap.Logger

func InitLg() *zap.Logger {
	var err error
	logger, err = zap.NewProduction() // Создание логгера
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	return logger
}

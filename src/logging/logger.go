package logging

import (
	"log"

	"go.uber.org/zap"
)

// create logger
func CreateLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}
	return logger
}

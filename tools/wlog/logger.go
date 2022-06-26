package wlog

import (
	"log"

	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger

func New() *zap.SugaredLogger {
	if sugar == nil {
		logger, err := zap.NewProduction()
		if err != nil {
			log.Fatal(err)
		}
		sugar = logger.Sugar()
	}

	return sugar
}

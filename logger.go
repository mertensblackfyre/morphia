

package main

import (
	"go.uber.org/zap"
)

var Sugar *zap.SugaredLogger

func Logger() {
	logger, _ := zap.NewProduction()
	Sugar = logger.Sugar()
}

func Sync() {
	if Sugar != nil {
		_ = Sugar.Sync()
	}
}


package main

import (
	"github.com/jkandasa/ytdl/cmd/app"
	"github.com/jkandasa/ytdl/pkg/types"
	"github.com/jkandasa/ytdl/pkg/version"
	"go.uber.org/zap"
)

func main() {
	cfg := types.Config{
		WebDirectory:    "/ui",
		EnableProfiling: false,
		BindAddress:     "0.0.0.0",
		Port:            8080,
	}

	// init logger
	loggerCfg := zap.NewDevelopmentConfig()
	logger, err := loggerCfg.Build()
	if err == nil {
		zap.ReplaceGlobals(logger)
	}

	// print version details
	zap.L().Info("version details", zap.Any("data", version.Get()))

	// start services
	app.Start(cfg)
}

package app

import (
	"fmt"
	"net/http"

	"github.com/jkandasa/ytdl/pkg/types"
	"go.uber.org/zap"
)

func Start(cfg types.Config) {
	hdlr, err := GetHandler(cfg)
	if err != nil {
		zap.L().Fatal("error on getting handler", zap.Error(err))
	}

	listenAddr := fmt.Sprintf("%s:%d", cfg.BindAddress, cfg.Port)
	zap.L().Info("listening...", zap.String("address", listenAddr))
	err = http.ListenAndServe(listenAddr, hdlr)
	if err != nil {
		zap.L().Fatal("error on starting handler", zap.Error(err))
	}
}

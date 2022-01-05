package utils

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

func WriteResponse(w http.ResponseWriter, data []byte) {
	_, err := w.Write(data)
	if err != nil {
		zap.L().Error("error on writing response", zap.Error(err))
		return
	}
}

func WriteJsonResponse(w http.ResponseWriter, data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	WriteResponse(w, bytes)
}

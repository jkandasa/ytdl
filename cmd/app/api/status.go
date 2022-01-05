package api

import (
	"net/http"

	"github.com/gorilla/mux"
	appUtils "github.com/jkandasa/ytdl/cmd/app/utils"
	"github.com/jkandasa/ytdl/pkg/version"
)

// RegisterStatusRoutes registers status,version api
func RegisterStatusRoutes(router *mux.Router) {
	router.HandleFunc("/api/version", versionData).Methods(http.MethodGet)
}

func versionData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	v := version.Get()
	appUtils.WriteJsonResponse(w, &v)
}

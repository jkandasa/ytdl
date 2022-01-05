package app

import (
	"net/http"

	"github.com/NYTimes/gziphandler"
	"github.com/gorilla/mux"
	handlerAPI "github.com/jkandasa/ytdl/cmd/app/api"
	"github.com/jkandasa/ytdl/pkg/types"

	appUtils "github.com/jkandasa/ytdl/cmd/app/utils"

	"github.com/rs/cors"
)

// GetHandler for http access
func GetHandler(cfg types.Config) (http.Handler, error) {
	router := mux.NewRouter()

	// register routes
	handlerAPI.RegisterStatusRoutes(router)
	handlerAPI.RegisterYoutubeRoutes(router)
	// Enable Profiling, if enabled
	if cfg.EnableProfiling {
		handlerAPI.RegisterPProfRoutes(router)
	}

	if cfg.WebDirectory != "" {
		fs := http.FileServer(http.Dir(cfg.WebDirectory))
		router.PathPrefix("/").Handler(fs)
	} else {
		defaultPage := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/plain")
			appUtils.WriteResponse(w, []byte("Web directory not configured."))
		}
		router.HandleFunc("/", defaultPage)
	}

	corsUtil := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		// Enable Debugging for testing, consider disabling in production
		Debug: false,
	})

	// Insert the middleware
	withCorsHandler := corsUtil.Handler(router)

	// update gzip handler
	withGzipHandler := gziphandler.GzipHandler(withCorsHandler)

	return withGzipHandler, nil
}

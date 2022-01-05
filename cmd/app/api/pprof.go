package api

import (
	"net/http"
	"net/http/pprof"

	"github.com/gorilla/mux"
)

// RegisterPProfRoutes registers pprof api
func RegisterPProfRoutes(router *mux.Router) {
	router.HandleFunc("/admin/debug/pprof/cmdline", pprof.Cmdline).Methods(http.MethodGet)
	router.HandleFunc("/admin/debug/pprof/profile", pprof.Profile).Methods(http.MethodGet)
	router.HandleFunc("/admin/debug/pprof/symbol", pprof.Symbol).Methods(http.MethodGet)
	router.HandleFunc("/admin/debug/pprof/trace", pprof.Trace).Methods(http.MethodGet)
	router.PathPrefix("/admin/debug/pprof/").HandlerFunc(pprof.Index).Methods(http.MethodGet)
}

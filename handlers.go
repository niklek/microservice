package main

import (
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/julienschmidt/httprouter"
	"github.com/niklek/microservice/internal/version"
	"github.com/sirupsen/logrus"
)

// Health live - liveness probe
func healthLive(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}

// Health ready - readiness probe
func healthReady(w http.ResponseWriter, _ *http.Request, _ httprouter.Params, isReady *atomic.Value) {
	if isReady == nil || !isReady.Load().(bool) {
		http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Custom not found
func notFound(w http.ResponseWriter, r *http.Request) {
	log.WithFields(logrus.Fields{
		"path":      r.URL.Path,
		"version":   version.Version,
		"commit":    version.Commit,
		"buildTime": version.BuildTime,
	}).Info("Requested unknown path")
	w.WriteHeader(http.StatusNotFound)
}

// Home handler
func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Request path:%s\n", r.URL.Path)
}

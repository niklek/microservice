package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/niklek/microservice/internal/version"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
	"sync/atomic"
)

var log = logrus.New()

func main() {
	log.SetFormatter(&logrus.JSONFormatter{})
	log.WithFields(logrus.Fields{
		"version":   version.Version,
		"commit":    version.Commit,
		"buildTime": version.BuildTime,
	}).Info("Starting service...")

	log.Info("Reading configuration...")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not set")
	}

	router := httprouter.New()
	router.GET("/", home)
	// liveness probe
	router.GET("/health/live", func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		w.WriteHeader(http.StatusOK)
	})
	// readiness probe
	isReady := &atomic.Value{}
	isReady.Store(false)
	go func() {
		// TODO prepare service
		time.Sleep(5 * time.Second)
		isReady.Store(true)
	}()
	router.GET("/health/ready", func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		if isReady == nil || !isReady.Load().(bool) {
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	log.Info("Service is ready to listen")
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal("Can not start the service:", err)
	}
}

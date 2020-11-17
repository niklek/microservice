package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/niklek/microservice/internal/version"
	"github.com/sirupsen/logrus"
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

	// channels for graceful shutdown
	interrupt := make(chan os.Signal, 1)
	shutdown := make(chan error, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	// routers
	router := httprouter.New()
	router.GET("/", home)
	router.GET("/health/live", healthLive)
	// a flag to indicate when the service is ready
	isReady := &atomic.Value{}
	isReady.Store(false)
	go func() {
		// TODO prepare the service
		time.Sleep(5 * time.Second)
		isReady.Store(true)
	}()
	router.GET("/health/ready", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		healthReady(w, r, p, isReady)
	})

	// server
	server := http.Server{
		Addr:    net.JoinHostPort("", port),
		Handler: router,
	}
	log.Info("Service is ready to listen")
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			shutdown <- err
		}
	}()

	// Lister for interruption and gracefully stop the server
	select {
	case x := <-interrupt:
		log.Info("Received signal", x.String())
	case err := <-shutdown:
		log.Error("Received an error from server", err)
	}
	log.Info("Stopping the service...")
	timeout, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	err := server.Shutdown(timeout)
	if err != nil {
		log.Error("Shutdown error", err)
	}
	log.Info("The service is stopped")
}

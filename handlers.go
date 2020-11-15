package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/niklek/microservice/internal/version"
	"github.com/sirupsen/logrus"
	"net/http"
)

// Home handler
func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.WithFields(logrus.Fields{
		"path":    r.URL.Path,
		"version": version.Version,
		"commit":  version.Commit,
	}).Info("Requested home")
	fmt.Fprintf(w, "Request path:%s\n", r.URL.Path)
}

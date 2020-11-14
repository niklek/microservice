package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.WithFields(logrus.Fields{
		"path": r.URL.Path,
	}).Info("Requested home")
	fmt.Fprintf(w, "Request path:%s\n", r.URL.Path)
}

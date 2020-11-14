package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	log.SetFormatter(&logrus.JSONFormatter{})
	log.Info("Starting service...")

	router := httprouter.New()
	router.GET("/", home)
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal("Can not start the service:", err)
	}
}

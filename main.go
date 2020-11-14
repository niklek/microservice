package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var log = logrus.New()

func main() {
	log.SetFormatter(&logrus.JSONFormatter{})
	log.Info("Starting service...")

	log.Info("Reading configuration...")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not set")
	}

	router := httprouter.New()
	router.GET("/", home)

	log.Info("Service is ready to listen")
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal("Can not start the service:", err)
	}
}

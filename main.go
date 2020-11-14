package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", home)
	http.ListenAndServe(":8000", router)
}
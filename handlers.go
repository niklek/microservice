package main

import (
	"net/http"
	"fmt"
	"github.com/julienschmidt/httprouter"
)

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Request path:%s", r.URL.Path)
}
package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Request path:%s\n", r.URL.Path)
}

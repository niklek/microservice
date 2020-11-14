package main

import (
	"net/http"
	"fmt"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Request path:%s", r.URL.Path)
}
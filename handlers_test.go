package main

import (
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test home path response
func TestHandlerHome(t *testing.T) {
	path := "/"
	respCodeExpected := http.StatusOK
	bodyExpected := "Request path:" + path + "\n"

	router := httprouter.New()
	router.GET(path, home)

	ts := httptest.NewServer(router)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/")
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Can not read response body: %s", err)
	}
	resp.Body.Close()

	if resp.StatusCode != respCodeExpected {
		t.Fatalf(
			"ResponseCode:%d, expected:%d",
			resp.StatusCode, respCodeExpected,
		)
	}

	if string(body) != bodyExpected {
		t.Fatalf(
			"Response body:%s, expected:%s",
			body, bodyExpected,
		)
	}
}

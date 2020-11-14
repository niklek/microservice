package main

import (
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerHome(t *testing.T) {
	path := "/"
	respCodeExpected := 200

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

	bodyExpected := "Request path:" + path + "\n"
	if string(body) != bodyExpected {
		t.Fatalf(
			"Response body:%s, expected:%s",
			body, bodyExpected,
		)
	}
}

package esa

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
)

// Stub テスト用のスタブ
func Stub(filename string, outRes interface{}) (*httptest.Server, *Client) {
	stub, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var statusCode int
		switch r.Method {
		case "GET":
			statusCode = http.StatusOK
		case "POST":
			statusCode = http.StatusCreated
		case "PATCH":
			statusCode = http.StatusOK
		case "DELETE":
			statusCode = http.StatusNoContent
		default:
			statusCode = http.StatusOK
		}
		w.WriteHeader(statusCode)
		w.Write([]byte(stub))
	}))
	c := NewClient("")
	c.baseURL = ts.URL

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}
	if err := json.Unmarshal([]byte(data), outRes); err != nil {
		log.Fatalln(err)
	}

	return ts, c
}

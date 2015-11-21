package esa
import (
	"net/http/httptest"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
)

func Stub(filename string, outRes interface{}) (*httptest.Server, *Client) {
	stub, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

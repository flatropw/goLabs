package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type Visitor struct {
	Host        string
	UserAgent   string
	RequestUri  string
	Headers     http.Header
	QueryParams url.Values
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		jsonVisitorData, err := json.MarshalIndent(
			Visitor{
				Host:        r.Host,
				UserAgent:   r.Header.Get("User-Agent"),
				RequestUri:  r.URL.Path,
				Headers:     r.Header,
				QueryParams: r.URL.Query(),
			}, "", "    ")
		if err != nil {
			panic(err)
		}
		_, err = w.Write(jsonVisitorData)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}

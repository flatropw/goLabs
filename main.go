package main

import (
	"fmt"
	"log"
	"net/http"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	_, _ = fmt.Fprint(w, "main page")
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, statusCode int) {
	w.WriteHeader(statusCode)
	_, _ = fmt.Fprintf(w, "%d", statusCode)
}

func main() {
	http.HandleFunc("/", mainHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

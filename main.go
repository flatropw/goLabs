package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const templateFolder = "template/"

var AllowedMethods = []string{
	http.MethodPost,
	http.MethodGet,
}

type ErrorData struct {
	StatusCode int
	Text       string
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAllowedMethod(r.Method) {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			_, _ = fmt.Fprint(w, err)
			return
		}

		address := r.FormValue("address")
		name := r.FormValue("name")

		tokenCookie := http.Cookie{
			Name:  "token",
			Value: name + ":" + address,
		}
		http.SetCookie(w, &tokenCookie)
	}

	http.ServeFile(w, r, templateFolder+"index.html")
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, statusCode int) {
	w.WriteHeader(statusCode)

	tmpl, _ := template.ParseFiles(templateFolder + "error.html")
	_ = tmpl.Execute(w, ErrorData{
		StatusCode: statusCode,
		Text:       http.StatusText(statusCode),
	})
}

func main() {
	http.HandleFunc("/", mainHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func IsAllowedMethod(method string) bool {
	for _, v := range AllowedMethods {
		if method == v {
			return true
		}
	}

	return false
}

package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Sadconf Platform 2021 555")
	})

	r.HandleFunc("/website/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("test.html")
		if err != nil {
			fmt.Println(err)
		}
		items := struct {
			Country string
			City    string
		}{
			Country: "Australia",
			City:    "Paris",
		}
		//	t.Execute(w, items)
		buf := &bytes.Buffer{}
		err = t.Execute(buf, items)
		if err != nil {
			// Send back error message, for example:
			http.Error(w, "Hey, Request was bad!", http.StatusBadRequest) // HTTP 400 status
		} else {
			// No error, send the content, HTTP 200 response status implied
			buf.WriteTo(w)
		}
		//	name := r.URL.Path[len("/greet/"):]
		//	fmt.Fprintf(w, "Hello %s\n", name)
	})

	http.ListenAndServe(":9990", r)
}

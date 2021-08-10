package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

type PageVariables struct {
	Date string
	Time string
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Sadconf Platform 2021 555111")
	})

	r.HandleFunc("/prueba/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Prueba")
	})
	r.HandleFunc("/website/", func(w http.ResponseWriter, r *http.Request) {

		ts, err := template.ParseFiles("test.html")
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error1111", 500)
			return
		}

		// We then use the Execute() method on the template set to write the template
		// content as the response body. The last parameter to Execute() represents any
		// dynamic data that we want to pass in, which for now we'll leave as nil.
		err = ts.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error222", 500)
		}
	})

	http.ListenAndServe(":9990", r)
}

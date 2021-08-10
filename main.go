package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Sadconf Platform 2021")
	})

	http.HandleFunc("/website/", func(w http.ResponseWriter, r *http.Request) {
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
		t.Execute(w, items)

		//	name := r.URL.Path[len("/greet/"):]
		//	fmt.Fprintf(w, "Hello %s\n", name)
	})

	http.ListenAndServe(":9990", nil)
}

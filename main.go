package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Sadconf Platform 2021 555")
	})

	r.HandleFunc("/website/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Sadconf Platform website")
	})

	http.ListenAndServe(":9990", r)
}

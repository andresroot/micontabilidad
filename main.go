package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type PageVariables struct {
	Date string
	Time string
}

func favicon(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s\n", r.RequestURI)
	w.Header().Set("Content-Type", "image/x-icon")
	w.Header().Set("Cache-Control", "public, max-age=7776000")
	fmt.Fprintln(w, "data:image/x-icon;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQEAYAAABPYyMiAAAABmJLR0T///////8JWPfcAAAACXBIWXMAAABIAAAASABGyWs+AAAAF0lEQVRIx2NgGAWjYBSMglEwCkbBSAcACBAAAeaR9cIAAAAASUVORK5CYII=\n")
}

func main() {
	r := mux.NewRouter()

	Raiz, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	if Raiz == "C:\\micontabilidad\\bin" {
		Raiz = ""
	} else {
		Raiz = Raiz + "/"
	}

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir(Raiz+"static/"))))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Sadconf Platform 2021 555111")
	})

	r.HandleFunc("/prueba/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Prueba")
	})

	// default
	r.HandleFunc("/default/", func(w http.ResponseWriter, r *http.Request) {

		ts, err := template.ParseFiles(Raiz + "vista/default.html")
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error1111"+err.Error()+Raiz, 500)
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

	r.HandleFunc("/website/", func(w http.ResponseWriter, r *http.Request) {

		//fmt.Println(dir)

		ts, err := template.ParseFiles(Raiz + "vista/test.html")
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error1111"+err.Error()+Raiz, 500)
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

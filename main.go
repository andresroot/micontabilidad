package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"time"
)

type PageVariables struct {
	Date string
	Time string
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Sadconf Platform 2021 555")
	})

	r.HandleFunc("/website/", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()              // find the time right now
		HomePageVars := PageVariables{ //store the date and time in a struct
			Date: now.Format("02-01-2006"),
			Time: now.Format("15:04:05"),
		}

		t, err := template.ParseFiles("test.html") //parse the html file homepage.html
		if err != nil {                            // if there is an error
			fmt.Fprintf(w, "template parsing error: "+err.Error()) // log it
		}
		err = t.Execute(w, HomePageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
		if err != nil {                  // if there is an error
			fmt.Fprintf(w, "template executing error: "+err.Error()) //log it
		}
	})

	http.ListenAndServe(":9990", r)
}

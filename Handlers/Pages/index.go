package app

import (
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	
	template, err := template.ParseFiles("../templates/pages/index.html")
	if err != nil {
		log.Fatal("error in page index")
	}
	err = template.Execute(w, nil)
	if err != nil {
		log.Fatal("error in executing template of index")
	}
}

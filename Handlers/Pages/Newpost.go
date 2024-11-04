package app

import (
	"html/template"
	"log"
	"net/http"
)

func Newpost(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("../templates/pages/Createpost.html")
	if err != nil {
		log.Fatal("error in page create post")
	}
	err = template.Execute(w, nil)
	if err != nil {
		log.Fatal("error in executing template of creating post")
	}
}

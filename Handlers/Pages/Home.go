package app

import (
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("username")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	template, err := template.ParseFiles("../templates/pages/home.html")
	if err != nil {
		log.Fatal("error in page home")
	}
	err = template.Execute(w, nil)
	if err != nil {
		log.Fatal("error in executing template of home")
	}
}

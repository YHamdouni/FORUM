package app

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	temp, _ := template.ParseFiles("../templates/pages/home.html")
	temp.Execute(w, "name.Value")
}

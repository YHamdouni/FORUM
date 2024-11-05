package app

import (
	"database/sql"
	"log"

	"net/http"
	"text/template"
)

var db *sql.DB

func SetDB(database *sql.DB) {
	db = database
}
func HandleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
	template, err := template.ParseFiles("../templates/pages/register.html")
	if err != nil {
		log.Fatal("error in page register")
	}
	name := r.FormValue("name")
	pass := r.FormValue("password")
	email := r.FormValue("email")
	if name == "" || pass == "" || email == "" {
		w.WriteHeader(http.StatusBadRequest)
		template.Execute(w, "please fill the form!")
		return
	}
	post := `INSERT INTO User (email,username,password_hash)
	VALUES (?, ?, ?)`
	_, err = db.Exec(post, email, name, pass)
	if err != nil {
		template.Execute(w, "email or name already exists")
	}
	http.Redirect(w, r, "/login", http.StatusMovedPermanently)
}
func Register(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("../templates/pages/register.html")
	if err != nil {
		log.Fatal("error in page register")
	}
	err = template.Execute(w, nil)
	if err != nil {
		log.Fatal("error in executing template of register")
	}
}

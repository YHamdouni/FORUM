package app

import (
	"database/sql"

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
	temp, _ := template.ParseFiles("../templates/pages/register.html")
	name := r.FormValue("name")
	pass := r.FormValue("password")
	email := r.FormValue("email")
	if name == "" || pass == "" || email == "" {
		w.WriteHeader(http.StatusBadRequest)
		temp.Execute(w, "please fill the form!")
		return
	}
	post := `INSERT INTO User (email,username,password_hash)
	VALUES (?, ?, ?)`
	_, err := db.Exec(post, email, name, pass)
	if err != nil {
		temp.Execute(w, "email or name already exists")
	}
	http.Redirect(w, r, "/login", http.StatusMovedPermanently)
}
func Register(w http.ResponseWriter, r *http.Request) {
	temp, _ := template.ParseFiles("../templates/pages/register.html")
	temp.Execute(w, nil)
}

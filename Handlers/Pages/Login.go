package app

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	template, err := template.ParseFiles("../templates/pages/login.html")
	if err != nil {
		log.Fatal("error in page login")
	}
	name := r.FormValue("name")
	pass := r.FormValue("password")

	if name == "" || pass == "" {
		w.WriteHeader(http.StatusBadRequest)
		template.Execute(w, "please fill the form!")
		return
	}
	post := `select password_hash,deja from user where username = ?`
	f := db.QueryRow(post, name)
	if f.Err() != nil {
		template.Execute(w, "username not found !")
		return
	}
	var passw string
	var deja int
	f.Scan(&passw, &deja)
	if pass != passw {
		template.Execute(w, "incorrect password !")
		return
	}
	c := http.Cookie{
		Name:     "username",
		Value:    name,
		Expires:  time.Now().Add(1 * time.Hour),
		Secure:   true,
		HttpOnly: true,
	}
	if deja == 0 {
		db.Exec(`UPDATE User SET deja = ? WHERE username = ?`, 1, name)
	} else {
		template.Execute(w, "walalala")
		return
	}
	http.SetCookie(w, &c)
	http.Redirect(w, r, "/home", http.StatusMovedPermanently)
}

func Login(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("../templates/pages/login.html")
	if err != nil {
		log.Fatal("error in page login")
	}
	err = template.Execute(w, nil)
	if err != nil {
		log.Fatal("error in executing template of login")
	}
}

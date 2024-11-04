package app

import (
	"html/template"
	"net/http"
	"time"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	temp, _ := template.ParseFiles("../templates/pages/login.html")
	name := r.FormValue("name")
	pass := r.FormValue("password")

	if name == "" || pass == "" {
		w.WriteHeader(http.StatusBadRequest)
		temp.Execute(w, "please fill the form!")
		return
	}
	post := `select password_hash,deja from user where username = ?`
	f := db.QueryRow(post, name)
	if f.Err() != nil {
		temp.Execute(w, "username not found !")
		return
	}
	var passw string
	var deja int
	f.Scan(&passw, &deja)
	if pass != passw {
		temp.Execute(w, "incorrect password !")
		return
	}
	c := http.Cookie{
		Name:     "username",
		Value:    name,
		Expires:  time.Now().Add(24 * time.Hour),
		Secure:   true,
		HttpOnly: true,
	}
	if deja == 0 {
		db.Exec(`UPDATE User SET deja = ? WHERE username = ?`, 1, name)
	} else {
		temp.Execute(w, "walalala")
		return
	}
	http.SetCookie(w, &c)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Login(w http.ResponseWriter, r *http.Request) {
	temp, _ := template.ParseFiles("../templates/pages/login.html")
	temp.Execute(w, nil)
}

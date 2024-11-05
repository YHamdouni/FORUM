package app

import (
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	name, _ := r.Cookie("username")
	c := http.Cookie{
		Name:   "username",
		MaxAge: -1,
	}
	http.SetCookie(w, &c)
	post := `select deja from user where username = ?`
	f := db.QueryRow(post, name.Value)
	var deja int
	f.Scan(&deja)
	if deja == 1 {
		db.Exec(`UPDATE User SET deja = ? WHERE username = ?`, 0, name.Value)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

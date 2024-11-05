package app

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Posts struct {
	Title   string
	Content string
}

func Newpost(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("username")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	template, err := template.ParseFiles("../templates/pages/Createpost.html")
	if err != nil {
		log.Fatal("error in page create post")
	}
	err = template.Execute(w, nil)
	if err != nil {
		log.Fatal("error in executing template of creating post")
	}
	// http.Redirect(w, r, "/home", 302)
}

func SubmitPost(w http.ResponseWriter, r *http.Request) {
	post := Posts{Title: r.FormValue("title"), Content: r.FormValue("content")}
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS Posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		text TEXT UNIQUE NOT NULL,
		username TEXT UNIQUE NOT NULL,
	);
	`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(post)

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

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
	usename := `INSERT FROM User(username)`
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS Posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE NOT NULL,
		title TEXT UNIQUE NOT NULL,
		content TEXT UNIQUE NOT NULL
	);
	`)
	p := `INSERT INTO Posts (username,title,content)
	VALUES (?, ?)`
	_, err = db.Exec(p, usename, post.Title, post.Content)
	if err != nil {
		log.Fatal("aji hnaya")
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(post)

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

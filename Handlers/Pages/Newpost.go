package app

import (
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
}

func SubmitPost(w http.ResponseWriter, r *http.Request) {
	post := Posts{Title: r.FormValue("title"), Content: r.FormValue("content")}
	cookie, err := r.Cookie("username")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	username := cookie.Value
	p := `INSERT INTO Posts (username, title, content) VALUES (?, ?, ?)`
	_, err = db.Exec(p, username, post.Title, post.Content)
	if err != nil {
		log.Fatal("Error inserting post: ", err)
	}

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

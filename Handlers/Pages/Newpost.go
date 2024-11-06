package app

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type Posts struct {
	Username string
	Title    string
	Category string
	Content  string
	Date     string
	Time     string
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
	now := time.Now()
	post := Posts{
		Title:    r.FormValue("title"),
		Category: r.FormValue("category"),
		Content:  r.FormValue("content"),
		Date:     now.Format("2006-01-02"),
		Time:     now.Format("15:04:05"),
	}
	cookie, err := r.Cookie("username")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	username := cookie.Value
	p := `INSERT INTO Posts (username, title, category, content, date, time) VALUES (?, ?, ?, ?, ?, ?)`
	_, err = db.Exec(p, username, post.Title, post.Category, post.Content, post.Date, post.Time)
	if err != nil {
		log.Fatal("Error inserting post: ", err)
	}

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

package app

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Username string
	Posts    []Posts
}

func Home(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	Username := cookie.Value

	rows, err := db.Query("SELECT title, content FROM Posts")
	if err != nil {
		log.Printf("Error fetching posts: %v", err)
		http.Error(w, "Could not load posts.", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var posts []Posts
	for rows.Next() {
		var post Posts
		if err := rows.Scan(&post.Title, &post.Content); err != nil {
			log.Printf("Error scanning posts: %v", err)
			http.Error(w, "Error loading posts.", http.StatusInternalServerError)
			return
		}
		posts = append(posts, post)
	}

	Data := PageData{
		Username: Username,
		Posts:    posts,
	}
	tmpl, err := template.ParseFiles("../templates/pages/home.html")
	if err != nil {
		log.Fatal("Error loading home page template: ", err)
	}
	err = tmpl.Execute(w, Data)
	if err != nil {
		log.Fatal("Error executing home page template: ", err)
	}
}

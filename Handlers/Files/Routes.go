package app

import (
	"database/sql"
	Pages "main/Handlers/Pages"
	"net/http"
)

func RegisterRoutes(database *sql.DB) {
	Pages.SetDB(database)
	http.HandleFunc("/styles/", StaticFiles)
	http.HandleFunc("/handleRegister", Pages.HandleRegister)
	http.HandleFunc("/handleLogin", Pages.HandleLogin)
	http.HandleFunc("/register", Pages.Register)
	http.HandleFunc("/login", Pages.Login)
	http.HandleFunc("/", Pages.Home)
	http.HandleFunc("/logout", Pages.Logout)
}

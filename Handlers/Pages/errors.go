package app

import (
	"html/template"
	"net/http"
)

type Err struct {
	StatusCode int
	Tittle     string
	Message    string
}

func ErrorHandler(w http.ResponseWriter, n int) {
	w.WriteHeader(n)
	template, err := template.ParseFiles("../templates/pages/error.html")
	if err != nil {
		http.Error(w, Errors[n].Tittle, n)
		return
	}
	err = template.Execute(w, Errors[n])
	if err != nil {
		http.Error(w, Errors[n].Tittle, n)
		return
	}
}

var Errors map[int]Err = map[int]Err{
	400: {
		StatusCode: 400,
		Tittle:     "Bad Request",
		Message:    "The server could not understand the request due to invalid syntax.",
	},
	404: {
		StatusCode: 404,
		Tittle:     "Page Not Found",
		Message:    "The page you are looking for might have been removed, had its name changed, or is temporarily unavailable.",
	},
	405: {
		StatusCode: 405,
		Tittle:     "Method Not Allowed",
		Message:    "The method specified in the request is not allowed for the resource identified by the request URI.",
	},
	500: {
		StatusCode: 500,
		Tittle:     "Internal Server Error",
		Message:    "The server encountered an unexpected condition that prevented it from fulfilling the request.",
	},
}

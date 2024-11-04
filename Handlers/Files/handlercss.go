package app

import (
	"net/http"
	"strings"
	Pages "main/Handlers/Pages"
)

func StaticFiles(w http.ResponseWriter, r *http.Request) {
	FileServer := http.StripPrefix("/styles/", http.FileServer(http.Dir("../../templates/styles/")))
	if strings.HasSuffix(r.URL.Path, "/") {
		Pages.ErrorHandler(w, http.StatusNotFound)
		return
	}
	FileServer.ServeHTTP(w, r)
}

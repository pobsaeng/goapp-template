package handlers

import (
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Name: "Example Page",
		Raw:  []string{"Content 1", "Content 2", "Content 3"},
		Path: "/",
	}
	RenderTemplateAndHandleError(w, "index.html", data)
}


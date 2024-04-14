package handlers

import (
	"net/http"
)

type SuccessData struct {
	Username string
	Role string
}

func HandleSuccess(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("Username")
	role := r.URL.Query().Get("Role")

	data := SuccessData{
		Username: username,
		Role: role,
	}
	RenderTemplateAndHandleError(w, "success.html", data)
}

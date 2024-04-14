package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Username string
	Role     string
 }

func HandleSignIn(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleSignInGet(w, r)
	case "POST":
		handleSignInPost(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleSignInGet(w http.ResponseWriter, _ *http.Request) {
	RenderTemplateAndHandleError(w, "signin.html", PageData{
		Name: "Sign in",
		Path: "/signin",
	})
}

func handleSignInPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("Error parsing form data:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	if username == "" || password == "" {
		RenderTemplateAndHandleError(w, "signin.html", PageData{
			Name:         "Sign in",
			Path:         "/signin",
			ErrorMessage: "Username and password are required.",
		})
		return
	}

	log.Printf("Username: %s, Password: %s", username, password)
	http.Redirect(w, r, fmt.Sprintf("/success?Username=%s&Role=%s", username, "Admin"), http.StatusSeeOther)

}
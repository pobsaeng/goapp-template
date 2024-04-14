package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

type PageData struct {
    Name string
    Raw  []string
    Content interface{}
    Path string
    ReloadPage bool
    ErrorMessage string
}

func renderTemplate(w http.ResponseWriter, tmplPath string, data interface{}) error {
    tmplFiles := []string{
        "templates/header.html",
        path.Join("pages", tmplPath),
        "templates/footer.html",
    }
    t, err := template.ParseFiles(tmplFiles...)
    if err != nil {
        return err
    }
    err = t.ExecuteTemplate(w, tmplPath, data)
    if err != nil {
        return err
    }
    return nil
}

func RenderTemplateAndHandleError(w http.ResponseWriter, tmpl string, data interface{}) {
	err := renderTemplate(w, tmpl, data)
	if err != nil {
		log.Println("Error rendering template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

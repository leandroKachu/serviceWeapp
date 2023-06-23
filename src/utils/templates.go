package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

func LoadTemplate() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

func RunTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)
}

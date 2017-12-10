package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

var templates *template.Template

//IndexHandler dada
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "./views/index.html", nil)
}

//RenderTemplate Template rendering function
func RenderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	t, err := template.ParseFiles(templateFile, "./views/partials/head.html", "./views/partials/footer.html")
	if err != nil {
		fmt.Fprintln(w, "not implemented yet !", err)
	}
	t.Execute(w, templateData)
}

func init() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

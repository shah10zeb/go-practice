package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// here we are reading from disk every single time on request (NOT efficient)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Printf("error in parsing template : %s , Error : %e", tmpl, err)
		return
	}
}
package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/template.html")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}

	err = parsedTemplate.Execute(w, data)
	if err != nil {
		fmt.Println("error executing template:", err)
	}
}

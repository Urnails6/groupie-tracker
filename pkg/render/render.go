package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// RenderTemplate renders an HTML template
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	parsedTemplate, err := template.ParseFiles("./templates/"+tmpl, "./templates/template.html")
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from error:", r)
		}
	}()

	err = parsedTemplate.Execute(w, data)
	if err != nil {
		fmt.Println("error executing template:", err)
	}
}

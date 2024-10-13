package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// here we are reading from disk every single time on request (NOT efficient)

func RenderTemplateOld(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Printf("error in parsing template : %s , Error : %e", tmpl, err)
		return
	}
}

//template cache
var tc= make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string){
	var tmpl *template.Template
	var err error

	//check if we already have in map

	_, inMap:= tc[t]

	if !inMap {
		//need to create template
		log.Println("Creating template and adding in Map")
		err= createTemplateCache(t)

		if err !=nil{
			log.Println(err)
		}

	}else{
		log.Printf("using Cache")
	}

	tmpl = tc[t]

	err= tmpl.Execute(w, nil)

	if err !=nil{
		log.Println(err)
	}

}

func createTemplateCache(t string) error{
templates:= []string{
	fmt.Sprintf("./templates/%s",t),
	"./templates/base.layout.tmpl",
}
//parse template
tmpl, err := template.ParseFiles(templates...)
if err!=nil{
	return err
}
tc[t]=tmpl
return nil
}
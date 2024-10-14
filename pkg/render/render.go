package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/shah10zeb/go-practice/pkg/config"
	"github.com/shah10zeb/go-practice/pkg/models"
)

var app *config.AppConfig

// set the config for template package
func NewTemplate(a *config.AppConfig) {
	app = a
}

// add Default Data
func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

// here we are reading from disk every single time on request (NOT efficient)

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	//get req template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Issue")
	}
	buf := new(bytes.Buffer)
	td = AddDefaultData(td)
	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	//render template

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// myCache:=make(map[string]*template.Template) OR
	myCache := map[string]*template.Template{}

	//get all of the files named *.page.tmpl

	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		log.Println(err)
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			log.Println(err)
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			log.Println(err)
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				log.Println(err)
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil

}

//***********************************************OLD WAY OF CACHING
//template cache
// var tc= make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string){
// 	var tmpl *template.Template
// 	var err error

// 	//check if we already have in map

// 	_, inMap:= tc[t]

// 	if !inMap {
// 		//need to create template
// 		log.Println("Creating template and adding in Map")
// 		err= createTemplateCache(t)

// 		if err !=nil{
// 			log.Println(err)
// 		}

// 	}else{
// 		log.Printf("using Cache")
// 	}

// 	tmpl = tc[t]

// 	err= tmpl.Execute(w, nil)

// 	if err !=nil{
// 		log.Println(err)
// 	}

// }

// func createTemplateCache(t string) error{
// templates:= []string{
// 	fmt.Sprintf("./templates/%s",t),
// 	"./templates/base.layout.tmpl",
// }
// //parse template
// tmpl, err := template.ParseFiles(templates...)
// if err!=nil{
// 	return err
// }
// tc[t]=tmpl
// return nil
// }

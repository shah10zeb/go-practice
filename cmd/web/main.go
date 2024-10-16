package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shah10zeb/go-practice/pkg/config"
	"github.com/shah10zeb/go-practice/pkg/handlers"
	"github.com/shah10zeb/go-practice/pkg/render"
)

const portNumber = ":8080"

// main application function
func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot Create template Cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)
	// http.HandleFunc("/about", handlers.Repo.About)
	// http.HandleFunc("/", handlers.Repo.Home)
	fmt.Printf("Starting port on %s", portNumber)
	// _ = http.ListenAndServe(portNumber, nil)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

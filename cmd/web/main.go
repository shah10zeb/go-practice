package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/shah10zeb/go-practice/pkg/config"
	"github.com/shah10zeb/go-practice/pkg/handlers"
	"github.com/shah10zeb/go-practice/pkg/render"
)

const portNumber = ":8080"

var app config.AppConfig

var session *scs.SessionManager

// main application function
func main() {

	session = scs.New()
	session.Lifetime = 24 * time.Hour

	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

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

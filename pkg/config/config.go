package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

//make sure config is not dependent on other local packages
//to avoid import cycle

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}

package config

import (
	"html/template"
	"log"
)

//make sure config is not dependent on other local packages
//to avoid import cycle

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}

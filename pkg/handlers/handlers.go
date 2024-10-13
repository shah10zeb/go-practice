package handlers

import (
	"net/http"

	"github.com/shah10zeb/go-practice/pkg/render"
)

// home handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "myhome.page.tmpl")
}

// about handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}

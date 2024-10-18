package handlers

import (
	"net/http"

	"github.com/shah10zeb/go-practice/pkg/config"
	"github.com/shah10zeb/go-practice/pkg/models"
	"github.com/shah10zeb/go-practice/pkg/render"
)

// Repository type
type Repository struct {
	App *config.AppConfig
}

// Repository variable
var Repo *Repository

// Creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//Newhandlers sets the repository for handlers

func NewHandlers(r *Repository) {
	Repo = r
}

// home handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "myhome.page.tmpl", &models.TemplateData{})
}

// about handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringmap := make(map[string]string)
	stringmap["test"] = "TEMPLATE DATA"
	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringmap["remote_ip"] = remoteIp
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{StringMap: stringmap})
}

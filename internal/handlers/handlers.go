package handlers

import (
	"github.com/Mooxtar/color-palette-explorer/internal/config"
	"github.com/Mooxtar/color-palette-explorer/internal/models"
	"github.com/Mooxtar/color-palette-explorer/internal/render"
	
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Most(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	render.RenderTemplate(w, r, "most.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Generate(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generate.page.tmpl", &models.TemplateData{})
}



package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/kguen/snippetbox/pkg/forms"
	"github.com/kguen/snippetbox/pkg/models"
)

type templateData struct {
	CurrentYear       int
	Snippet           *models.Snippet
	Snippets          []*models.Snippet
	Form              *forms.Form
	Flash             string
	AuthenticatedUser *models.User
	CSRFToken         string
}

func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	// Jan 2 15:04:05 2006 MST
	return t.UTC().Format("02 Jan 2006 at 15:04")
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	funcs := template.FuncMap{
		"humanDate": humanDate,
	}
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.html"))
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(funcs).ParseFiles(page)
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.html"))
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.html"))
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}
	return cache, nil
}

package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/huannguyen2114/golang-project/snippetbox/internal/models"
)

// Define a templateData type to act as the holding structure for
// any dymamic data that we want to pass to our HTML templates.
// At the moment it only contains one field, but we'll add more
// to it as the build progress

type templateData struct {
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
	CurrentYear int
	Form        any
	Flash       string
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

// Initialize a template.FuncMap object and store it in a blobal variable. This is
// essentially a string-keyed map which acts as a llokup between the names of our
// custom template funciotns and the functions themselves
var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	// Initialize a new map to act as the cache
	cache := map[string]*template.Template{}

	// Use the filepath.Glob() to get a slice of all filepaths that
	// match the pattern "./ui/html/pages/*.html". This will esssentially gives
	// us a slice of all the file paths for our application 'page' templates
	// like : [ui/html/pages/home.tmpl ui/html/pages/view.html]

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}
	// Loop through the page filepaths one-by one
	for _, page := range pages {

		name := filepath.Base(page)
		// The template.FuncMap must be registered with the template set before you

		// call the parseFiles() method. This means we have to use template.New() through
		// to create an empty template set, use the FUncs() method to register the
		// template.FuncMap, and then aprse the file as normal
		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.tmpl.html")
		if err != nil {
			return nil, err
		}

		// Call ParseGlob() * on this template set* to add any partials.
		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl.html")
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}
		// Add the template set to the map, using the name of the page as the key
		cache[name] = ts
	}
	return cache, nil
}

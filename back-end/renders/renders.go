package renders

import (
	"html/template"
	"path/filepath"
)

// CreateTemplateCache Create a Template Cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("../front-end/*.html")
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("../front-end/base.page.html")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err := ts.ParseGlob("../front-end/base.page.html")
			if err != nil {
				return myCache, err
			}
			myCache[name] = ts
		}
	}
	return myCache, nil
}

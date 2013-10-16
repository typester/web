package template

import (
	"path/filepath"
	"os"
	"strings"
	"net/http"
	"html/template"
)

var templateDir string
var tmpl *template.Template

func TemplateDir() string {
	return templateDir
}

func SetTemplateDir(dir string) (err error) {
	tmpl = template.New("web")
	
	err = filepath.Walk(dir, func (path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			name := strings.TrimPrefix(path, dir)
			if name[0] == os.PathSeparator {
				name = name[1:]
			}
			tmpl.New(name).ParseFiles(path)
		}
		return nil
	})

	return
}

func Render(w http.ResponseWriter, name string, args map[string]interface{}) error {
	return tmpl.ExecuteTemplate(w, name, args)
}





















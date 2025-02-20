package docs

import (
	"embed"
	httptemplate "html/template"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	apiFile   = "/static/swagger.yaml"
	indexFile = "template/index.tpl"
)

//go:embed static
var Static embed.FS

//go:embed template
var template embed.FS

func RegisterOpenAPIService(appName string, rtr *mux.Router) {
	rtr.Handle(apiFile, http.FileServer(http.FS(Static)))
	rtr.HandleFunc("/", handler(appName))
}

// handler returns an http handler that servers OpenAPI console for an OpenAPI spec at specURL.
func handler(title string) http.HandlerFunc {
	t, err := httptemplate.ParseFS(template, indexFile)
	if err != nil {
		panic(err)
	}

	return func(w http.ResponseWriter, req *http.Request) {
		err := t.Execute(w, struct {
			Title string
			URL   string
		}{
			title,
			apiFile,
		})
		if err != nil {
			panic(err)
		}
	}
}

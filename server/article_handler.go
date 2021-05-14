package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/cbroglie/mustache"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

type ArticleData struct {
	HTML      string
	CSS       string
	Title     string
	Date      string
	Location  string
	HeroImage string
}

func (s Server) ArticleHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.NotFound(w, r)
			return
		}

		articleData := &ArticleData{}
		jsonData, err := ioutil.ReadFile(filepath.Join(s.ResourcesDir, "submodules", "blog", "json", fmt.Sprint(r.FormValue("id"), ".json")))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(jsonData, articleData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		articleData.CSS = "blog-with-hero-image"
		if articleData.HeroImage == "" {
			articleData.CSS = "blog"
		}

		md := goldmark.New(
			goldmark.WithRendererOptions(html.WithUnsafe()),
			goldmark.WithExtensions(extension.Footnote, extension.Strikethrough, extension.Typographer),
		)

		mdData, err := ioutil.ReadFile(filepath.Join(s.ResourcesDir, "submodules", "blog", "posts", fmt.Sprint(r.FormValue("id"), ".md")))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		var htmlContent bytes.Buffer
		if err := md.Convert(mdData, &htmlContent); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		articleData.HTML = htmlContent.String()

		htmlData, err := mustache.RenderFile(filepath.Join(s.ResourcesDir, "mustache", "article.mustache"), articleData, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte(htmlData))
	}
}

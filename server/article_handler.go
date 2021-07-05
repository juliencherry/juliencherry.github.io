package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"

	"github.com/cbroglie/mustache"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

type ArticleData struct {
	HTML      string
	CSS       string
	ID        string
	Title     string
	Date      string
	Datetime  time.Time
	Location  string
	HeroImage string
}

func (s Server) ArticleHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.NotFound(w, r)
			return
		}

		articleData, err := Article(r.FormValue("id"), s.ResourcesDir)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		htmlData, err := mustache.RenderFile(filepath.Join(s.ResourcesDir, "mustache", "article.mustache"), articleData, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte(htmlData))
	}
}

func Article(id string, resourcesDir string) (*ArticleData, error) {
	articleData := &ArticleData{ID: id}

	jsonData, err := ioutil.ReadFile(filepath.Join(resourcesDir, "submodules", "blog", "json", fmt.Sprint(id, ".json")))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonData, articleData)
	if err != nil {
		return nil, err
	}

	articleData.CSS = "blog"

	articleData.Datetime, err = time.Parse("January 2, 2006", articleData.Date)
	if err != nil {
		return nil, err
	}

	md := goldmark.New(
		goldmark.WithRendererOptions(html.WithUnsafe()),
		goldmark.WithExtensions(extension.Footnote, extension.Strikethrough, extension.Typographer),
	)

	mdData, err := ioutil.ReadFile(filepath.Join(resourcesDir, "submodules", "blog", "posts", fmt.Sprint(id, ".md")))
	if err != nil {
		return nil, err
	}

	var htmlContent bytes.Buffer
	if err := md.Convert(mdData, &htmlContent); err != nil {
		return nil, err
	}
	articleData.HTML = htmlContent.String()

	return articleData, nil
}

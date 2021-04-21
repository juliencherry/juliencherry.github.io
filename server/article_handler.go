package server

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

func (s Server) ArticleHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		md := goldmark.New(
			goldmark.WithRendererOptions(html.WithUnsafe()),
			goldmark.WithExtensions(extension.Footnote, extension.Strikethrough, extension.Typographer),
		)

		mdData, err := ioutil.ReadFile(filepath.Join(s.ResourcesDir, "submodules", "blog", "posts", fmt.Sprint(r.FormValue("id"), ".md")))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		var htmlData bytes.Buffer
		if err := md.Convert(mdData, &htmlData); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		s.PageHandler("article", htmlData.String())(w, r)
	}
}

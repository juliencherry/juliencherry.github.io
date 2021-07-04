package server

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
	"sort"
	"strings"

	"github.com/cbroglie/mustache"
)

type BlogData struct {
	Posts     Posts
	HTML      string
	CSS       string
	Title     string
	Date      string
	Location  string
	HeroImage string
}

type Posts []*ArticleData

func (p Posts) Len() int {
	return len(p)
}

func (p Posts) Less(i, j int) bool {
	return p[i].Datetime.After(p[j].Datetime)
}

func (p Posts) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (s Server) BlogHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.NotFound(w, r)
			return
		}

		fileInfo, err := ioutil.ReadDir(filepath.Join(s.ResourcesDir, "submodules", "blog", "json"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		blogData := &BlogData{
			Title: "Blog",
			CSS:   "blog",
		}

		for _, file := range fileInfo {
			articleId := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
			articleData, err := Article(articleId, s.ResourcesDir)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			articleData.Date = articleData.Datetime.Format("2006-01-02")

			blogData.Posts = append(blogData.Posts, articleData)
		}

		sort.Sort(blogData.Posts)

		htmlData, err := mustache.RenderFile(filepath.Join(s.ResourcesDir, "mustache", "blog.mustache"), blogData, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte(htmlData))
	}
}

package server

import (
	"net/http"
	"path/filepath"
)

type Server struct {
	ResourcesDir string
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := http.NewServeMux()

	router.Handle("/", s.PageHandler("index"))
	router.Handle("/colophon", s.PageHandler("colophon"))
	router.Handle("/ethos", s.PageHandler("ethos"))
	router.Handle("/generate", s.PageHandler("generate"))
	router.Handle("/neurodiversity", s.PageHandler("neurodiversity"))
	router.Handle("/photography", s.PageHandler("photography"))
	router.Handle("/playlists", s.PageHandler("playlists"))

	router.Handle("/blog", s.BlogHandler())
	router.Handle("/article", s.ArticleHandler())

	router.Handle("/projects", s.PageHandler("projects"))
	router.Handle("/projects/chimerical-colors/", http.StripPrefix("/projects/chimerical-colors/", http.FileServer(http.Dir(filepath.Join(s.ResourcesDir, "submodules", "chimerical-colors")))))
	router.Handle("/projects/lambda-iota-engma/", http.StripPrefix("/projects/lambda-iota-engma/", http.FileServer(http.Dir(filepath.Join(s.ResourcesDir, "submodules", "lambda-iota-engma")))))

	router.Handle("/.well-known/", http.StripPrefix("/.well-known/", http.FileServer(http.Dir(filepath.Join(s.ResourcesDir, ".well-known")))))
	router.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir(filepath.Join(s.ResourcesDir, "css")))))
	router.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir(filepath.Join(s.ResourcesDir, "fonts")))))
	router.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir(filepath.Join(s.ResourcesDir, "images")))))
	router.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir(filepath.Join(s.ResourcesDir, "submodules", "blog", "img")))))
	router.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir(filepath.Join(s.ResourcesDir, "js")))))

	router.Handle("/favicon.ico", FileHandler(filepath.Join(s.ResourcesDir, "favicon.ico")))
	router.Handle("/robots.txt", FileHandler(filepath.Join(s.ResourcesDir, "robots.txt")))

	router.Handle("/github", RedirectHandler("https://github.com/juliencherry"))
	router.Handle("/linkedin", RedirectHandler("https://www.linkedin.com/in/juliencherry/"))

	router.ServeHTTP(w, r)
}

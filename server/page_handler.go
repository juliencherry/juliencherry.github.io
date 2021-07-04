package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/cbroglie/mustache"
)

type PageData struct {
	HTML               string
	CSS                string
	Title              string
	Date               string
	Photos             []Photo
	Playlists          []Playlist
	Projects           []Project
	StatesAndProvinces []StatesAndProvinces
}

type Photo struct {
	Title        string
	Src          string
	Alt          string
	Location     string
	Date         string
	Aperture     string
	ShutterSpeed string
	Lens         string
	Film         string
	Notes        string
}

type Playlist struct {
	Name        string
	Description string
	Embed       string
}

type Project struct {
	Name        string
	Link        string
	Description string
}

type StatesAndProvinces struct {
	Name         string
	Abbreviation string
	Capital      string
}

func (s Server) PageHandler(name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.NotFound(w, r)
			return
		}

		jsonData, err := ioutil.ReadFile(filepath.Join(s.ResourcesDir, "json", fmt.Sprint(name, ".json")))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		pageData := &PageData{}
		err = json.Unmarshal(jsonData, pageData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		htmlData, err := mustache.RenderFile(filepath.Join(s.ResourcesDir, "mustache", fmt.Sprint(name, ".mustache")), pageData, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte(htmlData))
	}
}

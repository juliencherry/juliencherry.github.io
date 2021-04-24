package server

import "net/http"

func RedirectHandler(URL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, URL, http.StatusFound)
	}
}

package server

import "net/http"

func FileHandler(filepath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.NotFound(w, r)
			return
		}

		http.ServeFile(w, r, filepath)
	}
}

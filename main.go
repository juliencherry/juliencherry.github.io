package main

import (
	"log"
	"net/http"
	"path/filepath"

	"juliencherry.net/sass"
	"juliencherry.net/server"
)

func main() {
	resourcesDir := "resources"
	sass := sass.Compiler{resourcesDir}

	dirs := []string{
		"",
		filepath.Join("submodules", "chimerical-colors"),
		filepath.Join("submodules", "lambda-iota-engma"),
	}

	for _, dir := range dirs {
		err := sass.Compile(dir)
		if err != nil {
			log.Fatal(err)
		}
	}

	port := ":8081"
	log.Printf("Starting server on %s\n", port)
	if err := http.ListenAndServe(port, server.Server{resourcesDir}); err != nil {
		log.Fatal(err)
	}
}

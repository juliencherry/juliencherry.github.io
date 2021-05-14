package sass

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/wellington/go-libsass"
)

type Compiler struct {
	ResourcesDir string
}

func (c Compiler) Compile(baseDir string) error {
	baseDir = filepath.Join(c.ResourcesDir, baseDir)
	srcDir := filepath.Join(baseDir, "scss")
	outDir := filepath.Join(baseDir, "css")

	dirEntries, err := os.ReadDir(srcDir)
	if err != nil {
		return err
	}

	for _, dirEntry := range dirEntries {
		srcFilename := dirEntry.Name()
		if srcFilename[0] == '_' {
			continue
		}

		r, err := os.Open(filepath.Join(srcDir, srcFilename))
		if err != nil {
			return err
		}

		var b bytes.Buffer
		comp, err := libsass.New(&b, r)
		if err != nil {
			return err
		}

		err = comp.Option(libsass.IncludePaths([]string{srcDir}))
		if err != nil {
			return err
		}

		err = comp.Run()
		if err != nil {
			return err
		}

		err = os.MkdirAll(outDir, os.ModePerm)
		if err != nil {
			return err
		}

		stylesheet := strings.TrimSuffix(srcFilename, filepath.Ext(srcFilename))
		err = ioutil.WriteFile(filepath.Join(outDir, fmt.Sprint(stylesheet, ".css")), b.Bytes(), os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

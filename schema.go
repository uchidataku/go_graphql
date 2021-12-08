package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func Dirwalk(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read dir: %w", err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			// Recursively calls Dirwalk in the case of a directory
			p, err := Dirwalk(filepath.Join(dir, file.Name()))
			if err != nil {
				return nil, fmt.Errorf("dirwalk %s: %w", filepath.Join(dir, file.Name()), err)
			}
			// Merge into the caller's "paths" variable.
			paths = append(paths, p...)
			continue
		}
		// Now that we've reached a leaf (file) in the directory tree, we'll add it to "paths" variable.
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths, nil
}

func main() {
	paths, err := Dirwalk("./schema")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("schema.graphql")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for _, path := range paths {
		s, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		f.Write(s)
		f.WriteString("\n")
	}
}

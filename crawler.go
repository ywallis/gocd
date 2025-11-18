package main

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"strings"
)

func crawl_dir(dir *string) {

	hashes := make(map[string][]file)

	err := filepath.WalkDir(*dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %q: %v\n", path, err)
			return err
		}

		if d.IsDir() {
			return nil
		}
		fileName := d.Name()

		// Break flow for dotfiles
		if strings.HasPrefix(fileName, ".") {
			return nil
		}
		fileHash, err := hash_file(path)

		if err != nil {
			fmt.Printf("Error hashing file: %s\n", fileName)
			return err
		}

		hashes[fileHash] = append(hashes[fileHash], file{name: fileName, path: path})

		return nil
	})

	if err != nil {
		log.Fatalf("Error walking the path: %v\n", err)
	}
	for _, files := range hashes {
		if len(files) > 1 {
			fmt.Println("Found conflict:")
			smallestPath := files[0]
			for _, f := range files {
				fmt.Printf("%s at %s\n", f.name, f.path)
				if len(f.path) < len(smallestPath.path) {
					smallestPath = f
				}
			}
			fmt.Printf("Will keep: %s\n", smallestPath.name)
			fmt.Println()
		}
	}
}

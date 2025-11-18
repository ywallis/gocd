package main

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
)

func crawl_dir(dir *string) {

	files := make(map[string]bool)
	var dupes []string

	err := filepath.WalkDir(*dir, func(path string, d fs.DirEntry, err error) error {
		// 1. Handle errors
		if err != nil {
			fmt.Printf("Error accessing path %q: %v\n", path, err)
			return err // Stop walking this path
		}

		// 2. Print the path
		fileName := d.Name()

		if _, ok := files[fileName]; ok {
			dupes = append(dupes, fileName)
		} else {
			files[fileName] = true
		}

		// 3. Continue walking
		return nil
	})

	if err != nil {
		log.Fatalf("Error walking the path: %v\n", err)
	}
	fmt.Println("Found duplicate file names:", dupes)
}

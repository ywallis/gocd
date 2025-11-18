package main

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
)

func crawl_dir(dir *string) {

	hashes := make(map[string]bool)
	var dupes []string

	err := filepath.WalkDir(*dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %q: %v\n", path, err)
			return err
		}

		if d.IsDir() {
			return nil
		}
		fileName := d.Name()
		fileHash, err := hash_file(path)

		if err != nil {
			fmt.Printf("Error hashing file: %s\n", fileName)
			return err
		}

		if _, ok := hashes[fileHash]; ok {
			dupes = append(dupes, fileName)
		} else {
			hashes[fileHash] = true
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Error walking the path: %v\n", err)
	}
	fmt.Println("Found duplicate files:")
	for _, file := range dupes {
		fmt.Printf("- %s\n", file)
	}
}

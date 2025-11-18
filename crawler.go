package main

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"strings"
)

func crawlDir(dir *string) {

	hashes := make(map[string][]file)
	emptyDirs := []string{}

	err := filepath.WalkDir(*dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %q: %v\n", path, err)
			return err
		}

		fileName := d.Name()

		if d.IsDir() {
			if strings.HasPrefix(fileName, ".") {
				return filepath.SkipDir
			}
			isEmpty, err := isDirEmpty(path)
			if err != nil {
				return err
			}
			if isEmpty {
				emptyDirs = append(emptyDirs, path)

			}
			return nil
		}

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
	listConflits(hashes)

	if len(emptyDirs) > 0 {
		fmt.Println("Found some empty directories:")
		for _, dir := range emptyDirs {
			fmt.Printf("-%s\n", dir)
		}
	}
}

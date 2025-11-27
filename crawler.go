package main

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"strings"
)

func crawlDir(dir *string, limit int) {

	emptyDirs := []string{}

	sem := make(chan struct{}, limit)

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
			processDirectory(path, sem)
			return nil
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Error walking the path: %v\n", err)
	}

	if len(emptyDirs) > 0 {
		fmt.Println("Found some empty directories:")
		for _, dir := range emptyDirs {
			fmt.Printf("-%s\n", dir)
		}
	}
}

package main

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"strings"
	"sync"
)

func crawlDir(dir *string, limit int) {

	hashes := make(map[string][]file)
	emptyDirs := []string{}

	var wg sync.WaitGroup
	var mu sync.Mutex

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
			return nil
		}

		// Break flow for dotfiles
		if strings.HasPrefix(fileName, ".") {
			return nil
		}
		wg.Add(1)
		go func(p, name string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			fileHash, err := hash_file(path)
			if err != nil {
				fmt.Printf("Error hashing file: %s\n", fileName)
				return
			}
			mu.Lock()
			hashes[fileHash] = append(hashes[fileHash], file{name: fileName, path: path})
			mu.Unlock()
		}(path, fileName)

		return nil
	})

	if err != nil {
		log.Fatalf("Error walking the path: %v\n", err)
	}
	wg.Wait()
	listConflits(hashes)

	if len(emptyDirs) > 0 {
		fmt.Println("Found some empty directories:")
		for _, dir := range emptyDirs {
			fmt.Printf("-%s\n", dir)
		}
	}
}

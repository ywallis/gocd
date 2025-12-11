package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

func loadTypes() (map[string]string, error) {

	types := make(map[string]string)
	bytes, err := os.ReadFile("extensions.json")
	if err != nil {
		return types, err
	}
	err = json.Unmarshal(bytes, &types)
	if err != nil {
		return types, err
	}
	return types, nil
}

func fileExists(path string) bool {

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		return true
	}
}

func findNextPath(path string) string {
	fileComponents := strings.Split(path, ".")
	basePath := fileComponents[0]
	extension := fileComponents[len(fileComponents)-1]

	i := 1

	for {
		nStr := strconv.Itoa(i)
		pathWithNum := fmt.Sprintf("%s_%s.%s", basePath, nStr, extension)
		fmt.Println(pathWithNum)
		if fileExists(pathWithNum) {
			i++
		} else {
			return pathWithNum
		}
	}
}
func isDirEmpty(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err
}

func listConflits(hashes map[string][]file) {

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

func processDirectory(dirPath string, sem chan struct{}) error {

	hashes := make(map[string][]file)
	var wg sync.WaitGroup
	var mu sync.Mutex

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}

	for _, entry := range entries {

		fileName := entry.Name()
		fullPath := filepath.Join(dirPath, fileName)

		if entry.IsDir() || strings.HasPrefix(fileName, ".") {
			continue
		}
		wg.Add(1)
		go func(p, name string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			fileHash, err := hash_file(p)
			if err != nil {
				fmt.Printf("Error hashing file: %s\n", fileName)
				return
			}
			mu.Lock()
			hashes[fileHash] = append(hashes[fileHash], file{name: fileName, path: fullPath})
			mu.Unlock()
		}(fullPath, fileName)
	}

	wg.Wait()
	listConflits(hashes)
	return nil
}

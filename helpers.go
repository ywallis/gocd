package main

import (
	"fmt"
	"io"
	"os"
)

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

package main

import (
	"os"
	"path/filepath"
	"strings"
)

func organize(dirPath string) error {

	types, err := loadTypes()
	if err != nil {
		return err
	}
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		fileName := entry.Name()
		fileComponents := strings.Split(fileName, ".")
		extension := fileComponents[len(fileComponents)-1]
		if entry.IsDir() || strings.HasPrefix(fileName, ".") {
			continue
		}
		fileType, exist := types[extension]
		if !exist {
			fileType = "Other"
		}
		typeDirPath := filepath.Join(dirPath, fileType)
		err := os.MkdirAll(typeDirPath, 0777)
		if err != nil {
			return err
		}
		currentFilePath := filepath.Join(dirPath, fileName)
		newFilePath := filepath.Join(dirPath, fileType, fileName)

		if fileExists(newFilePath) {
			newFilePath = findNextPath(newFilePath)
			os.Rename(currentFilePath, newFilePath)
		} else {
			os.Rename(currentFilePath, newFilePath)
		}

	}

	return nil
}

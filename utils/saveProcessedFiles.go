package utils

import (
	"os"
	"strings"
)

func createDirectoryIfNotExists(dirname string) (err error) {
	_, err = os.Stat(dirname)

	if os.IsNotExist(err) {
		err = os.MkdirAll(dirname, 0755)
	}

	return err
}

func SaveProcessedFiles(files map[string][]byte, destPath string) (savedFileCount int, err error) {
	err = createDirectoryIfNotExists(destPath)

	if err != nil {
		return 0, err
	}

	for fullFileName, fileData := range files {

		lastDirectoryIndex := strings.LastIndex(fullFileName, "\\")

		fileName := fullFileName[lastDirectoryIndex+1:]

		newFullFileName := destPath + "\\" + fileName

		if err != nil {
			return 0, err
		}

		file, err := os.Create(newFullFileName)

		if err != nil {
			return 0, err
		}

		defer file.Close()

		_, err = file.Write(fileData)

		if err != nil {
			return 0, err
		}

		savedFileCount++
	}

	return savedFileCount, nil
}

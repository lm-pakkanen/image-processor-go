package utils

import (
	"fmt"
	"os"

	"github.com/zecuel/go/image-processor/helpers"
)

func getFullFilePath(pathOptions *helpers.PathOptions, fileName string, origin string) string {
	if origin == "src" {
		return fmt.Sprintf("%s\\%s", pathOptions.Src, fileName)
	}

	if origin == "dest" {
		return fmt.Sprintf("%s\\%s", pathOptions.Dest, fileName)
	}

	return ""
}

func isFileAlreadyProcessed(pathOptions *helpers.PathOptions, fileName string) bool {
	fullFileName := getFullFilePath(pathOptions, fileName, "dest")
	_, err := os.Stat(fullFileName)

	// If error is nil, file already exists
	return err == nil
}

func getFilePaths(pathOptions *helpers.PathOptions) ([]string, error) {
	file, err := os.Open(pathOptions.Src)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	fileNames, err := file.Readdirnames(0)

	if err != nil {
		return nil, err
	}

	return fileNames, nil
}

func GetFiles(pathOptions *helpers.PathOptions) (files helpers.FileMap, foundFilesCount int, err error) {
	files = make(helpers.FileMap, 100)

	fileNames, err := getFilePaths(pathOptions)

	if err != nil {
		return nil, foundFilesCount, err
	}

	for _, fileName := range fileNames {
		sFileName := helpers.String(fileName)

		// Unsupported file format, skip fetching and don't increment found counter
		if !sFileName.EndsWith(".jpg") && !sFileName.EndsWith(".png") {
			continue
		}

		if isFileAlreadyProcessed(pathOptions, fileName) {
			foundFilesCount++
			continue
		}

		fullName := getFullFilePath(pathOptions, fileName, "src")

		data, err := os.ReadFile(fullName)

		if err != nil {
			return nil, foundFilesCount, err
		}

		files[fullName] = data
		foundFilesCount++
	}

	return files, foundFilesCount, nil
}

package utils

import (
	"fmt"
	"os"
)

func fileAlreadyProcessed(fileName string, dest string) bool {
	fullFileName := fmt.Sprintf("%s\\%s", dest, fileName)
	_, err := os.Stat(fullFileName)

	// If error is nil, file already exists
	return err == nil
}

func getFullFilePath(src string, fileName string) string {
	return fmt.Sprintf("%s\\%s", src, fileName)
}

func getFilePaths(src string) ([]string, error) {
	file, err := os.Open(src)

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

func GetFiles(src string, dest string) (fileData map[string][]byte, foundFilesCount int, err error) {
	fileNames, err := getFilePaths(src)

	if err != nil {
		return nil, foundFilesCount, err
	}

	var fileDatas map[string][]byte = make(map[string][]byte, 100)

	for _, fileName := range fileNames {
		// TODO: check for .jpg or .png ending

		if fileAlreadyProcessed(fileName, dest) {
			foundFilesCount++
			continue
		}

		fullName := getFullFilePath(src, fileName)

		data, err := os.ReadFile(fullName)

		if err != nil {
			return nil, foundFilesCount, err
		}

		fileDatas[fullName] = data
		foundFilesCount++
	}

	return fileDatas, foundFilesCount, nil
}

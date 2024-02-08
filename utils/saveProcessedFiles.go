package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/zecuel/go/image-processor/helpers"
)

func createDirectoryIfNotExists(dirName string) (err error) {
	_, err = os.Stat(dirName)

	if os.IsNotExist(err) {
		err = os.MkdirAll(dirName, 0755)
	}

	return err
}

func SaveProcessedFiles(files helpers.FileMap, pathOptions *helpers.PathOptions) (savedFileCount int, err error) {
	err = createDirectoryIfNotExists(pathOptions.Dest)

	if err != nil {
		return 0, err
	}

	for fullFileName, fileData := range files {

		lastDirectoryIndex := strings.LastIndex(fullFileName, "\\")

		fileName := fullFileName[lastDirectoryIndex+1:]
		destinationFullFileName := fmt.Sprintf("%s\\%s", pathOptions.Dest, fileName)

		if err != nil {
			return 0, err
		}

		file, err := os.Create(destinationFullFileName)

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

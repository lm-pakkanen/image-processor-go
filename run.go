package main

import (
	"flag"
	"fmt"

	"github.com/zecuel/go/image-processor/utils"
)

const ROOT_SOURCE_PATH string = "C:\\Users\\Zecuel\\Desktop\\"

const DEFAULT_SOURCE_PATH string = "Legnum\\Pix\\Raw"
const DEFAULT_DESTINATION_PATH string = "Legnum\\Pix\\Processed"
const DEFAULT_MAX_DIMENSION_PX int = 800
const DEFAULT_QUALITY int = 75

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	srcPathFlagPtr := flag.String("src", DEFAULT_SOURCE_PATH, "source of images")
	destPathFlagPtr := flag.String("dest", DEFAULT_DESTINATION_PATH, "destination for images")
	maxDimPxFlagPtr := flag.Int("maxdimpx", DEFAULT_MAX_DIMENSION_PX, "maximum dimension of image in pixels")
	imgQualityFlagPtr := flag.Int("quality", DEFAULT_QUALITY, "quality of image 1-100")

	flag.Parse()

	srcPath := ROOT_SOURCE_PATH + *srcPathFlagPtr
	destPath := ROOT_SOURCE_PATH + *destPathFlagPtr

	var maxDimPx int = *maxDimPxFlagPtr
	var imgQuality int = *imgQualityFlagPtr

	// -> 25 - 100
	if imgQuality > 100 {
		imgQuality = 100
	} else if imgQuality < 25 {
		imgQuality = 25
	}

	fmt.Printf("Starting image processor with settings:\n - source path: %s\n - destination path: %s\n\n", srcPath, destPath)

	var filesToProcessCount int = 0
	var processedCount int = 0

	defer func() {
		fmt.Printf("\nProcessed %d/%d files\n", processedCount, filesToProcessCount)
	}()

	fileDatas, foundFilesCount, err := utils.GetFiles(srcPath, destPath)
	panicIfErr(err)

	filesToProcessCount = foundFilesCount

	processedFiles, err := utils.ProcessFiles(fileDatas, maxDimPx, imgQuality)
	panicIfErr(err)

	savedFileCount, err := utils.SaveProcessedFiles(processedFiles, destPath)
	panicIfErr(err)

	processedCount = savedFileCount

}

package main

import (
	"flag"
	"fmt"

	"github.com/lm-pakkanen/image-processor-go/helpers"
	"github.com/lm-pakkanen/image-processor-go/utils"
)

const ROOT_SOURCE_PATH string = "C:\\Users\\Zecuel\\Desktop\\"

const DEFAULT_SOURCE_PATH string = "Legnum\\Pix\\Raw"
const DEFAULT_DESTINATION_PATH string = "Legnum\\Pix\\Processed"
const DEFAULT_MAX_DIMENSION_PX int = 800
const DEFAULT_QUALITY int = 75

func main() {
	var filesToProcessCount int = 0
	var processedCount int = 0

	srcPathFlagPtr := flag.String("src", DEFAULT_SOURCE_PATH, "source of images")
	destPathFlagPtr := flag.String("dest", DEFAULT_DESTINATION_PATH, "destination for images")
	maxDimPxFlagPtr := flag.Int("maxdimpx", DEFAULT_MAX_DIMENSION_PX, "maximum dimension of image in pixels")
	imgQualityFlagPtr := flag.Int("quality", DEFAULT_QUALITY, "quality of image 1-100")

	flag.Parse()

	srcPath := ROOT_SOURCE_PATH + *srcPathFlagPtr
	destPath := ROOT_SOURCE_PATH + *destPathFlagPtr

	var maxDimPx int = *maxDimPxFlagPtr
	var imgQuality int = *imgQualityFlagPtr

	// 25 <= imgQuality <= 100
	if imgQuality > 100 {
		imgQuality = 100
	} else if imgQuality < 25 {
		imgQuality = 25
	}

	pathOptions := &helpers.PathOptions{Src: srcPath, Dest: destPath}
	processorOptions := &helpers.ProcessorOptions{MaximumDimensionPx: maxDimPx, ImgQuality: imgQuality}

	fmt.Printf(
		"Starting image processor with settings:\n"+
			" - source path: %s\n"+" - destination path: %s\n"+
			" - maximum dimension: %dpx\n"+
			" - image quality: %d/100\n\n",
		pathOptions.Src, pathOptions.Dest, processorOptions.MaximumDimensionPx, processorOptions.ImgQuality)

	defer func() {
		fmt.Printf("\nProcessed %d/%d files\n", processedCount, filesToProcessCount)
	}()

	fileDatas, filesToProcessCount, err := utils.GetFiles(pathOptions)
	helpers.PanicIfErr(err)

	processedFiles, err := utils.ProcessFiles(fileDatas, processorOptions)
	helpers.PanicIfErr(err)

	savedFileCount, err := utils.SaveProcessedFiles(processedFiles, pathOptions)
	helpers.PanicIfErr(err)

	processedCount = savedFileCount
}

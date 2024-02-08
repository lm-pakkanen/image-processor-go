package utils

import (
	"bytes"
	"errors"
	"image"
	"image/jpeg"
	"image/png"

	"github.com/lm-pakkanen/image-processor-go/helpers"
	"github.com/nfnt/resize"
)

func getNewDimension(imageInfo image.Config, options *helpers.ProcessorOptions) (heightPx int, widthPx int) {
	var imageHeightPx int = imageInfo.Height
	var imageWidthPx int = imageInfo.Width

	var maxDimensionPx int = 0

	if imageHeightPx > imageWidthPx {
		maxDimensionPx = imageHeightPx
	} else {
		maxDimensionPx = imageWidthPx
	}

	var dimensionMultiplier float64 = float64(options.MaximumDimensionPx) / float64(maxDimensionPx)

	// Image is already smaller than maximumDimension
	if dimensionMultiplier > 1.0 {
		return imageHeightPx, imageWidthPx
	}

	var newHeightPx int = int(float64(imageHeightPx) * dimensionMultiplier)
	var newWidthPx int = int(float64(imageWidthPx) * dimensionMultiplier)

	return newHeightPx, newWidthPx
}

func imgToBytes(img image.Image, format string, options *helpers.ProcessorOptions) (imgBytes []byte, err error) {
	fileDataBuf := new(bytes.Buffer)

	if format == "png" {
		pngEncoder := &png.Encoder{CompressionLevel: png.DefaultCompression}
		err = pngEncoder.Encode(fileDataBuf, img)
	} else if format == "jpg" {
		err = jpeg.Encode(fileDataBuf, img, &jpeg.Options{Quality: options.ImgQuality})
	} else {
		return nil, errors.New("UNSUPPORTED FILE FORMAT")
	}

	if err != nil {
		return nil, err
	}

	imgBytes = fileDataBuf.Bytes()
	return imgBytes, nil
}

func ProcessFiles(files helpers.FileMap, options *helpers.ProcessorOptions) (resultFiles helpers.FileMap, err error) {
	resultFiles = make(helpers.FileMap, 100)

	for fileName, fileData := range files {

		var imageFile image.Image
		var imageInfo image.Config

		var err1 error
		var err2 error

		var format string

		sFileName := helpers.String(fileName)

		if sFileName.EndsWith(".png") {
			format = "png"
			imageFile, err1 = png.Decode(bytes.NewReader(fileData))
			imageInfo, err2 = png.DecodeConfig(bytes.NewReader(fileData))
		} else if sFileName.EndsWith(".jpg") {
			format = "jpg"
			imageFile, _, err = image.Decode(bytes.NewReader(fileData))
			imageInfo, _, err2 = image.DecodeConfig(bytes.NewReader(fileData))
		} else {
			return nil, errors.New("UNSUPPORTED FILE FORMAT")
		}

		if err1 != nil {
			return nil, err
		}

		if err2 != nil {
			return nil, err2
		}

		heightPx, widthPx := getNewDimension(imageInfo, options)

		resizedImageFile := resize.Resize(uint(widthPx), uint(heightPx), imageFile, resize.Lanczos3)
		imgAsBytes, err := imgToBytes(resizedImageFile, format, options)

		if err != nil {
			return nil, err
		}

		resultFiles[fileName] = imgAsBytes

	}

	return resultFiles, nil
}

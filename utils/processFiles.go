package utils

import (
	"bytes"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"strings"

	"github.com/nfnt/resize"
)

func getNewDimension(imageInfo image.Config, maximumDimensionPx int) (height int, width int) {
	var imageHeight int = imageInfo.Height
	var imageWidth int = imageInfo.Width

	var maxCurrDimension int = 0

	if imageHeight > imageWidth {
		maxCurrDimension = imageHeight
	} else {
		maxCurrDimension = imageWidth
	}

	var dimensionMultiplier float64 = float64(maximumDimensionPx) / float64(maxCurrDimension)

	// Image is already smaller than maximumDimension
	if dimensionMultiplier > 1.0 {
		return imageHeight, imageWidth
	}

	var newHeight int = int(float64(imageHeight) * dimensionMultiplier)
	var newWidth int = int(float64(imageWidth) * dimensionMultiplier)

	return newHeight, newWidth
}

func imgToBytes(img image.Image, imgQuality int, format string) (imgBytes []byte, err error) {
	fileDataBuf := new(bytes.Buffer)

	if format == "png" {
		pngEncoder := &png.Encoder{CompressionLevel: png.DefaultCompression}
		err = pngEncoder.Encode(fileDataBuf, img)
	} else if format == "jpg" {
		err = jpeg.Encode(fileDataBuf, img, &jpeg.Options{Quality: imgQuality})
	} else {
		return nil, errors.New("UNSUPPORTED FILE FORMAT")
	}

	if err != nil {
		return nil, err
	}

	imgBytes = fileDataBuf.Bytes()
	return imgBytes, nil
}

func ProcessFiles(fileDatas map[string][]byte, maximumDimensionPx int, imgQuality int) (resultFileDatas map[string][]byte, err error) {
	resultFileDatas = make(map[string][]byte, 100)

	for fileName, fileData := range fileDatas {

		var imageFile image.Image
		var imageInfo image.Config

		var err1 error
		var err2 error

		var format string

		if strings.HasSuffix(fileName, ".png") {
			format = "png"
			imageFile, err1 = png.Decode(bytes.NewReader(fileData))
			imageInfo, err2 = png.DecodeConfig(bytes.NewReader(fileData))
		} else if strings.HasSuffix(fileName, ".jpg") {
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

		resizeheight, resizeWidth := getNewDimension(imageInfo, maximumDimensionPx)

		resizedImageFile := resize.Resize(uint(resizeWidth), uint(resizeheight), imageFile, resize.Lanczos3)

		imgAsBytes, err := imgToBytes(resizedImageFile, imgQuality, format)

		if err != nil {
			return nil, err
		}

		resultFileDatas[fileName] = imgAsBytes

	}

	return resultFileDatas, nil
}

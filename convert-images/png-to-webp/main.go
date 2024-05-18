package main

import (
	"errors"
	"image"
	"image/png"
	"os"

	"golang.org/x/image/webp"
)

func convertToWebp(sourcePath string, outputPath string, quality int) error {
	// Open the original image
	f, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		return err
	}

	// Encode the image as WebP with desired quality
	var webpData []byte
	switch img.ColorModel() {
	case png.CM_RGBA, png.CM_RGB:
		webpData, err = webp.Encode(webp.RGBAQualityEncoder{Quality: quality}, img)
	default:
		return errors.New("unsupported image color model")
	}
	if err != nil {
		return err
	}

	// Write the WebP data to the output file
	out, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = out.Write(webpData)
	return err
}

func main() {
	err := convertToWebp("image.png", "image.webp", 80) // Adjust quality (0-100)
	if err != nil {
		// handle error
	}
}

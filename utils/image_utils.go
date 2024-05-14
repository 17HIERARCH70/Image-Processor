package utils

import (
	"fmt"
	"github.com/h2non/bimg"
	"os"
	"path/filepath"
	"strings"
)

func ProcessFolder(folderPath string, quality, size int, blurRadius uint32, format, outputDirectory, nameSuffix string) {
	// Get list of files in folder
	files, err := os.ReadDir(folderPath)
	if err != nil {
		fmt.Println("Error reading folder:", err)
		return
	}

	// Process each file in folder
	for _, file := range files {
		// Check if file is an image
		if IsImage(file.Name()) {
			// Full path to the file
			filePath := filepath.Join(folderPath, file.Name())
			// Process image
			ProcessImage(filePath, quality, size, blurRadius, format, outputDirectory, nameSuffix)
		}
	}
}

func ProcessImage(filePath string, quality, size int, blurRadius uint32, format, outputDirectory, nameSuffix string) {
	// Read image
	buffer, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Initialize bimg image
	img := bimg.NewImage(buffer)

	// Apply blur to image
	if blurRadius > 0 {
		options := bimg.Options{
			GaussianBlur: bimg.GaussianBlur{Sigma: float64(blurRadius)},
		}
		newImage, err := img.Process(options)
		if err != nil {
			fmt.Println("Error applying blur:", err)
			return
		}
		img = bimg.NewImage(newImage)
	}

	// Resize image if size is specified
	if size > 0 {
		options := bimg.Options{
			Width: size * 1024,
		}
		newImage, err := img.Process(options)
		if err != nil {
			fmt.Println("Error resizing image:", err)
			return
		}
		img = bimg.NewImage(newImage)
	}

	// Set output format
	var imageType bimg.ImageType
	switch strings.ToLower(format) {
	case "jpeg", "jpg":
		imageType = bimg.JPEG
	case "png":
		imageType = bimg.PNG
	case "webp":
		imageType = bimg.WEBP
	default:
		imageType = imgType(img)
	}

	// Generate output file name
	outputFileName := generateOutputFileName(filePath, nameSuffix, blurRadius, size, imageType)
	outputFilePath := filepath.Join(outputDirectory, outputFileName)

	// Save image
	options := bimg.Options{
		Quality: quality,
		Type:    imageType,
	}
	newImage, err := img.Process(options)
	if err != nil {
		fmt.Println("Error processing image:", err)
		return
	}

	err = bimg.Write(outputFilePath, newImage)
	if err != nil {
		fmt.Println("Error writing image:", err)
		return
	}

	fmt.Println("Image processed:", filePath)
}

func generateOutputFileName(filePath, nameSuffix string, blurRadius uint32, size int, imageType bimg.ImageType) string {
	baseName := strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
	var suffix string
	if blurRadius > 0 && size > 0 {
		suffix = "_blurred_and_compressed"
	} else if blurRadius > 0 {
		suffix = "_blurred"
	} else if size > 0 {
		suffix = "_compressed"
	}
	if nameSuffix != "" {
		suffix = "_" + nameSuffix
	}

	ext := filepath.Ext(filePath)
	switch imageType {
	case bimg.JPEG:
		ext = ".jpg"
	case bimg.PNG:
		ext = ".png"
	case bimg.WEBP:
		ext = ".webp"
	}

	return baseName + suffix + ext
}

func IsImage(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".webp" || ext == ".gif"
}

func imgType(img *bimg.Image) bimg.ImageType {
	switch img.Type() {
	case "jpeg":
		return bimg.JPEG
	case "png":
		return bimg.PNG
	case "webp":
		return bimg.WEBP
	default:
		return bimg.UNKNOWN
	}
}

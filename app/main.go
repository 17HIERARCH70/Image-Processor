package main

import (
	"fmt"
	"github.com/17HIERARCH70/Image-Processor/utils"
	"os"

	"github.com/spf13/cobra"
)

var (
	folderPath      string
	quality         int
	size            int
	blurRadius      uint
	format          string
	outputDirectory string
	nameSuffix      string
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "image-processor",
		Short: "Image Processor is a tool to process images in various ways",
		Run: func(cmd *cobra.Command, args []string) {
			// Create output directory if it doesn't exist
			if _, err := os.Stat(outputDirectory); os.IsNotExist(err) {
				if err := os.MkdirAll(outputDirectory, os.ModePerm); err != nil {
					fmt.Println("Error creating output directory:", err)
					return
				}
			}

			// Process folder
			utils.ProcessFolder(folderPath, quality, size, uint32(blurRadius), format, outputDirectory, nameSuffix)
		},
	}

	rootCmd.Flags().StringVarP(&folderPath, "folder", "f", "img", "Path to the folder containing images")
	rootCmd.Flags().IntVarP(&quality, "quality", "q", 100, "Quality of encode of images (0-100)")
	rootCmd.Flags().IntVarP(&size, "size", "s", 0, "Size of each photo will not be more than this size (in MB)")
	rootCmd.Flags().UintVarP(&blurRadius, "blur", "b", 0, "Radius for Box blur")
	rootCmd.Flags().StringVarP(&format, "format", "F", "", "Format all photos into special format (png, jpg, jpeg, webp)")
	rootCmd.Flags().StringVarP(&outputDirectory, "output", "o", "img_out", "Output directory for processed images")
	rootCmd.Flags().StringVarP(&nameSuffix, "name", "n", "", "Specific suffix after name of image (for example compressed_and_blurred)")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

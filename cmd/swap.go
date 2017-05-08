package cmd

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"  // needed for image.Decode
	_ "image/png" // needed for voldemort.png
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/anthonynsimon/bild/blend"
	"github.com/nfnt/resize"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
)

// swapCmd represents the swap command
var swapCmd = &cobra.Command{
	Use:   "swap",
	Short: "Face Swap",
	Long:  `If a match is found in the image, do a faceswap`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(imageURL) == 0 {
			fmt.Println("Use --image to specify an image URL to identify")
		} else {
			identifyResponse := identify(imageURL)
			swap(imageURL, identifyResponse)
		}
	},
}

func init() {
	RootCmd.AddCommand(swapCmd)
	swapCmd.Flags().StringVarP(&imageURL, "image", "i", "", "image to identify")
}

func swap(imageURL string, identifyResponse IdentifyResponse) {
	// Get the image from the URL
	request := gorequest.New()
	response, _, err := request.Get(imageURL).End()
	if err != nil {
		panic(err)
	}

	// Get filename from URL
	srcURL, _ := url.Parse(imageURL)
	pathParts := strings.Split(srcURL.Path, "/")
	filename := pathParts[len(pathParts)-1]
	if strings.HasSuffix(strings.ToLower(filename), ".jpg") == false {
		// Add .jpg extension if the original image doesn't have it
		filename = filename + ".jpg"
	}

	background, _, decodeErr := image.Decode(response.Body)
	if decodeErr != nil {
		panic(decodeErr)
	}

	// Initialize a canvas.
	canvas := image.NewRGBA(background.Bounds())
	draw.Draw(canvas, canvas.Bounds(), background, image.ZP, draw.Src)

	// Replace each matched face
	for _, face := range identifyResponse.Faces {
		if face.Matched {
			// Load replacement image into newFace
			replacementImage, err := os.Open(filepath.Join("swap.png"))
			if err != nil {
				panic(err)
			}
			defer replacementImage.Close()
			newFace, _, err := image.Decode(replacementImage)
			if err != nil {
				panic(err)
			}

			// Draw the newFace where it was found on top of the original image
			drawPoint := image.Point{face.Rect.Left, face.Rect.Top}
			oldFaceSize := image.Point{face.Rect.Width, face.Rect.Height}

			fmt.Printf("Face of size %v identified at %v \n", oldFaceSize, drawPoint)

			// Resize newFace to fit the dimensions of the old face
			newFace = resize.Resize(
				uint(face.Rect.Width),
				uint(face.Rect.Height),
				newFace,
				resize.Bilinear)
			newFaceRect := newFace.Bounds()
			rect := image.Rectangle{drawPoint, drawPoint.Add(oldFaceSize)}
			draw.Draw(canvas, rect, newFace, newFaceRect.Min, draw.Src)
		}
	}

	// Blend the original image with the new image
	result := blend.Normal(background, canvas)

	// Output the new image
	outputImage, outputErr := os.Create(filename)
	if outputErr != nil {
		panic(err)
	}
	defer outputImage.Close()
	jpeg.Encode(outputImage, result, &jpeg.Options{Quality: jpeg.DefaultQuality})

}

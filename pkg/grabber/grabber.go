package grabber

import (
	"fmt"

	opencv "github.com/go-opencv/go-opencv/opencv"
	app "github.com/ryanbmilbourne/syr-sudoku-backend/pkg"
)

// GrabPuzzle parses a puzzle structure from a provided image
func GrabPuzzle(fileName string) (app.PuzzleState, error) {

	// Load the image.
	fmt.Println("LOADING THE FUCKING IMAGE")
	baseImg := opencv.LoadImage(fileName, 0)
	if baseImg == nil {
		return app.PuzzleState{}, fmt.Errorf("Opencv: Failed to load image: %v", fileName)
	}
	defer baseImg.Release()
	fmt.Println("LOADED THE FUCKING IMAGE")

	// Some pre-processing.
	// Gaussian blur the original image very slightly.  This removes some noise.
	fmt.Println("BLUR THAT SHIT")
	blurredImg := baseImg.Clone()
	defer blurredImg.Release()

	kernelSize := 11
	sigma := float64(0.3)*float64(float64(kernelSize-1)*float64(0.5)-1) + float64(0.8)
	opencv.Smooth(baseImg, blurredImg, opencv.CV_GAUSSIAN, kernelSize, kernelSize, sigma, 0)
	fmt.Println("SO BLURRY")

	fmt.Println("SAVE THE BLUR")
	opencv.SaveImage(fileName+"-blurred.jpg", blurredImg, nil)
	fmt.Println("WE GOT IT FOREVUH")

	// Now that it's blurred, do the needful.

	return app.PuzzleState{}, nil
}

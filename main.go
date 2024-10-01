package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strings"
	"time"
)

// Configuration constants
const (
	ProgressBarLength = 100
	ColorConfigFile   = ".year_progress_colors.json"
)

// ColorPalette contains named colors loaded from JSON or defaults
var ColorPalette map[string]string

// Default colors to use if JSON loading fails
var defaultColors = map[string]string{
	"Red":    "\033[31m",
	"Green":  "\033[32m",
	"Blue":   "\033[34m",
	"Yellow": "\033[33m",
	"Reset":  "\033[0m",
}

// loadColors reads the color definitions from a JSON file
func loadColors(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("Warning: Failed to open color config file: %v", err)
		log.Println("Using default colors")
		ColorPalette = defaultColors
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Warning: Failed to read color config file: %v", err)
		log.Println("Using default colors")
		ColorPalette = defaultColors
		return
	}

	err = json.Unmarshal(bytes, &ColorPalette)
	if err != nil {
		log.Printf("Warning: Failed to parse color config: %v", err)
		log.Println("Using default colors")
		ColorPalette = defaultColors
		return
	}

	// Ensure that we always have a Reset color
	if _, exists := ColorPalette["Reset"]; !exists {
		ColorPalette["Reset"] = "\033[0m"
	}
}

// calculateYearProgress returns the percentage of the year completed
func calculateYearProgress() float64 {
	now := time.Now()
	yearStart := time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, time.UTC)
	nextYearStart := yearStart.AddDate(1, 0, 0)

	elapsed := time.Since(yearStart)
	total := nextYearStart.Sub(yearStart)

	return (float64(elapsed) / float64(total)) * 100
}

// randomColor returns a random color from the ColorPalette
func randomColor() string {
	colors := make([]string, 0, len(ColorPalette))
	for name, code := range ColorPalette {
		if name != "Reset" {
			colors = append(colors, code)
		}
	}

	max := big.NewInt(int64(len(colors)))
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		log.Printf("Warning: Failed to generate random number: %v", err)
		return ColorPalette["Reset"] // Return Reset color as fallback
	}

	return colors[n.Int64()]
}

// renderProgressBar generates a colored progress bar
func renderProgressBar(percentage float64, length int) string {
	filledLength := int(percentage / 100 * float64(length))
	emptyLength := length - filledLength

	filledColor := randomColor()
	emptyColor := randomColor()

	filled := strings.Repeat(filledColor+"█"+ColorPalette["Reset"], filledLength)
	empty := strings.Repeat(emptyColor+"░"+ColorPalette["Reset"], emptyLength)

	return fmt.Sprintf("[%s%s]", filled, empty)
}

func main() {
	loadColors(ColorConfigFile)

	progress := calculateYearProgress()
	progressBar := renderProgressBar(progress, ProgressBarLength)

	fmt.Printf("Year Progress is: %.2f%%\n", progress)
	fmt.Println(progressBar)
}

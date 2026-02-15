package main

import (
	"crypto/rand"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Version information
const (
	Version           = "1.0.0"
	ProgressBarLength = 30
	DefaultConfigFile = ".year_progress_colors.json"
)

// ColorPalette contains named colors loaded from JSON or defaults
var ColorPalette map[string]string

// Default colors to use if JSON loading fails
var defaultColors = map[string]string{
	"Red":    "\033[31m",
	"Green":  "\033[32m",
	"Blue":   "\033[34m",
	"Yellow": "\033[33m",
	"Cyan":   "\033[36m",
	"Magenta": "\033[35m",
	"Reset":  "\033[0m",
}

// loadColors reads the color definitions from a JSON file
func loadColors(configPath string) error {
	file, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &ColorPalette)
	if err != nil {
		return err
	}

	// Ensure that we always have a Reset color
	if _, exists := ColorPalette["Reset"]; !exists {
		ColorPalette["Reset"] = "\033[0m"
	}

	return nil
}

// getConfigPath returns the path to the color config file
func getConfigPath() string {
	// Check XDG_CONFIG_HOME first
	if xdgConfig := os.Getenv("XDG_CONFIG_HOME"); xdgConfig != "" {
		configPath := filepath.Join(xdgConfig, "year-progress", DefaultConfigFile)
		if _, err := os.Stat(configPath); err == nil {
			return configPath
		}
	}

	// Check home directory
	homeDir, err := os.UserHomeDir()
	if err == nil {
		configPath := filepath.Join(homeDir, DefaultConfigFile)
		if _, err := os.Stat(configPath); err == nil {
			return configPath
		}
	}

	// Fallback to current directory
	return DefaultConfigFile
}

// calculateYearProgress returns the percentage of the year completed
func calculateYearProgress(t time.Time) float64 {
	yearStart := time.Date(t.Year(), time.January, 1, 0, 0, 0, 0, t.Location())
	nextYearStart := yearStart.AddDate(1, 0, 0)

	elapsed := t.Sub(yearStart)
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

	if len(colors) == 0 {
		return defaultColors["Reset"]
	}

	max := big.NewInt(int64(len(colors)))
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		log.Printf("Warning: Failed to generate random number: %v", err)
		return defaultColors["Reset"]
	}

	return colors[n.Int64()]
}

// renderProgressBar generates a colored progress bar
func renderProgressBar(percentage float64, length int) string {
	if percentage < 0 {
		percentage = 0
	}
	if percentage > 100 {
		percentage = 100
	}

	filledLength := int(percentage / 100 * float64(length))
	emptyLength := length - filledLength

	filledColor := randomColor()
	emptyColor := randomColor()

	filled := strings.Repeat(filledColor+"█"+ColorPalette["Reset"], filledLength)
	empty := strings.Repeat(emptyColor+"░"+ColorPalette["Reset"], emptyLength)

	return fmt.Sprintf("[%s%s]", filled, empty)
}

func main() {
	// Parse flags
	version := flag.Bool("v", false, "Show version")
	versionLong := flag.Bool("version", false, "Show version")
	forceColor := flag.Bool("force-color", false, "Force color output")
	configPath := flag.String("config", "", "Path to config file")
	flag.Parse()

	if *version || *versionLong {
		fmt.Printf("year-progress version %s\n", Version)
		os.Exit(0)
	}

	// Initialize color palette with defaults
	ColorPalette = make(map[string]string)
	for k, v := range defaultColors {
		ColorPalette[k] = v
	}

	// Load colors from config file
	configFile := *configPath
	if configFile == "" {
		configFile = getConfigPath()
	}

	if err := loadColors(configFile); err != nil {
		// Silently use defaults if config file not found
		// Only log if there's a real error (not "file not found")
		if !os.IsNotExist(err) {
			log.Printf("Warning: Failed to load color config: %v", err)
			log.Println("Using default colors")
		}
	}

	// Get progress for current time
	progress := calculateYearProgress(time.Now())
	progressBar := renderProgressBar(progress, ProgressBarLength)

	// Disable colors if not a TTY and --force-color not set
	if !*forceColor {
		// Colors are already raw ANSI, no special handling needed
		// This is a simplified approach - for production use a color library
	}

	fmt.Printf("Year Progress: %.2f%% %s\n", progress, progressBar)
}

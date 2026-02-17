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
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Version information
const (
	Version           = "1.0.3"
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

// hexColorRegex matches hex color codes like #FF0000 or FF0000
var hexColorRegex = regexp.MustCompile(`^#?([0-9A-Fa-f]{6})$`)

// hexToANSI converts a hex color code to ANSI 256-color escape sequence
func hexToANSI(hex string) string {
	// Remove # if present
	hex = strings.TrimPrefix(hex, "#")

	// Parse hex color
	r, err := strconv.ParseInt(hex[0:2], 16, 64)
	if err != nil {
		return ""
	}
	g, err := strconv.ParseInt(hex[2:4], 16, 64)
	if err != nil {
		return ""
	}
	b, err := strconv.ParseInt(hex[4:6], 16, 64)
	if err != nil {
		return ""
	}

	// Convert to ANSI 256 color
	// Formula for true color to 256 color
	if r == int64(b) && int64(b) == g {
		// Grayscale
		if r < 8 {
			return "\033[232m"
		}
		if r > 248 {
			return "\033[231m"
		}
		gray := int((r - 8) / 247 * 24)
		return fmt.Sprintf("\033[38;5;%dm", gray)
	}

	// Color
	rIdx := int(r / 51)
	gIdx := int(g / 51)
	bIdx := int(b / 51)

	ansi := 16 + (rIdx * 36) + (gIdx * 6) + bIdx
	return fmt.Sprintf("\033[38;5;%dm", ansi)
}

// processColor processes a color string - converts hex to ANSI if needed
func processColor(color string) string {
	// Check if it's a hex color
	if hexColorRegex.MatchString(color) {
		ansi := hexToANSI(color)
		if ansi != "" {
			return ansi
		}
	}

	// Return as-is (assume it's already an ANSI code)
	return color
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

	// First, parse raw JSON to detect hex colors
	var rawColors map[string]interface{}
	err = json.Unmarshal(bytes, &rawColors)
	if err != nil {
		return err
	}

	// Process colors: convert hex to ANSI
	ColorPalette = make(map[string]string)
	for name, value := range rawColors {
		colorStr, ok := value.(string)
		if !ok {
			continue
		}
		ColorPalette[name] = processColor(colorStr)
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
	percentageOnly := flag.Bool("percentage", false, "Output only the percentage number")
	jsonOutput := flag.Bool("json", false, "Output in JSON format")
	configPath := flag.String("config", "", "Path to config file")
	flag.Parse()

	if *version || *versionLong {
		fmt.Printf("year-progress version %s\n", Version)
		os.Exit(0)
	}

	// Get progress for current time
	progress := calculateYearProgress(time.Now())

	// Output formats
	if *percentageOnly {
		fmt.Printf("%.2f\n", progress)
		return
	}

	// Initialize color palette with defaults (only needed for progress bar)
	ColorPalette = make(map[string]string)
	for k, v := range defaultColors {
		ColorPalette[k] = v
	}

	// Load colors from config file if provided (silent - no logging)
	configFile := *configPath
	if configFile == "" {
		configFile = getConfigPath()
	}

	// Try to load - silently use defaults if file doesn't exist or can't be read
	loadColors(configFile)

	progressBar := renderProgressBar(progress, ProgressBarLength)

	if *jsonOutput {
		// For JSON, output clean percentage only (no ANSI codes)
		fmt.Printf("{\"percentage\":%.2f}\n", progress)
	} else {
		fmt.Printf("Year Progress: %.2f%% %s\n", progress, progressBar)
	}
}

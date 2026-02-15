package main

import (
	"strings"
	"testing"
	"time"
)

func TestCalculateYearProgress(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		minPct   float64
		maxPct   float64
	}{
		{
			name:   "January 1st should be 0%",
			input:  time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
			minPct: 0,
			maxPct: 0.1,
		},
		{
			name:   "July 1st should be around 50%",
			input:  time.Date(2024, time.July, 1, 12, 0, 0, 0, time.UTC),
			minPct: 49,
			maxPct: 51,
		},
		{
			name:   "December 31st should be ~100%",
			input:  time.Date(2024, time.December, 31, 23, 59, 59, 0, time.UTC),
			minPct: 99.9,
			maxPct: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateYearProgress(tt.input)
			if result < tt.minPct || result > tt.maxPct {
				t.Errorf("calculateYearProgress() = %v, want between %v and %v", result, tt.minPct, tt.maxPct)
			}
		})
	}
}

func TestRenderProgressBar(t *testing.T) {
	tests := []struct {
		name        string
		percentage  float64
		length      int
		wantFilled  int // expected filled characters (not including empty)
	}{
		{
			name:       "0 percent",
			percentage: 0,
			length:     10,
			wantFilled: 0,
		},
		{
			name:       "50 percent",
			percentage: 50,
			length:     10,
			wantFilled: 5,
		},
		{
			name:       "100 percent",
			percentage: 100,
			length:     10,
			wantFilled: 10,
		},
		{
			name:       "negative percentage clamped to 0",
			percentage: -10,
			length:     10,
			wantFilled: 0,
		},
		{
			name:       "over 100 clamped to 100",
			percentage: 150,
			length:     10,
			wantFilled: 10,
		},
	}

	// Initialize ColorPalette for tests
	ColorPalette = defaultColors

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := renderProgressBar(tt.percentage, tt.length)
			// Check that it starts with '[' and ends with ']'
			if result[0] != '[' || result[len(result)-1] != ']' {
				t.Errorf("renderProgressBar() should be wrapped in brackets, got %q", result)
			}
			// Count filled characters (█)
			filled := strings.Count(result, "█")
			if filled != tt.wantFilled {
				t.Errorf("renderProgressBar() filled = %v, want %v, result = %q", filled, tt.wantFilled, result)
			}
		})
	}
}

func TestLoadColors(t *testing.T) {
	// Test loading non-existent file
	ColorPalette = nil
	err := loadColors("/nonexistent/path/colors.json")
	if err == nil {
		t.Error("Expected error loading non-existent file")
	}
}

func TestGetConfigPath(t *testing.T) {
	// This test just verifies the function runs without panic
	// The actual path depends on the system
	path := getConfigPath()
	if path == "" {
		t.Error("getConfigPath() returned empty string")
	}
}

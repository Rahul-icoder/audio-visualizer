package visualizer

import (
	"fmt"
	"os"
	"golang.org/x/term"
)

// RenderFrequencies renders the frequency bars on the terminal and centers the content
func RenderFrequencies(frequencies []float64) {
	maxHeight := 20 // Default maximum height of the bars
	numBars := 40   // Number of bars

	// Get current tmux pane size
	width, height, err := getTerminalSize()
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		return
	}

	// Adjust maxHeight to fit within the terminal pane height
	if height < maxHeight {
		maxHeight = height - 2 // Leave some space for the terminal prompt
	}

	// Normalize and scale the frequencies
	scaledHeights := make([]int, numBars)
	for i, freq := range frequencies[:numBars] {
		height := int(freq * float64(maxHeight))
		if height > maxHeight {
			height = maxHeight
		}
		scaledHeights[i] = height
	}

	// Calculate the position to center the content horizontally
	offsetX := (width - numBars*2) / 2 // The width of each bar is 2 (due to spacing)

	// Ensure bars fit vertically and center the bars in the available space
	offsetY := height - maxHeight - 1 // Start from the bottom of the screen

	// If the terminal is too small, adjust the offsetY to avoid negative values
	if offsetY < 0 {
		offsetY = 0
	}

	// Clear the screen and move the cursor to the top-left corner
	fmt.Print("\033[H\033[2J") // Clear the screen

	// Render bars from the bottom upwards
	for row := maxHeight; row >= 1; row-- { // Loop from maxHeight down to 1
		// Move the cursor to the correct position for the current row
		fmt.Printf("\033[%d;%dH", offsetY+(maxHeight-row), offsetX)

		// Print the bars for the current row
		for _, height := range scaledHeights {
			if height >= row {
				// Color the bar based on the height
				colorCode := getColorCode(height)
				fmt.Printf("\033[%dm█ \033[0m", colorCode) // Print filled bar with color
			} else {
				fmt.Print("  ") // Print empty space
			}
		}
	}

	// Reset cursor to the top to avoid scrolling (move cursor back to top-left)
	fmt.Print("\033[H")
}

// getColorCode returns a color code based on the height of the bar
func getColorCode(height int) int {
	// Example of color scaling: green for low, yellow for medium, red for high
	if height <= 8 {
		return 32 // Green
	} else if height <= 16 {
		return 33 // Yellow
	} else {
		return 31 // Red
	}
}

// getTerminalSize gets the current terminal width and height
func getTerminalSize() (int, int, error) {
	// Get the size of the terminal window
	fd := int(os.Stdout.Fd())
	width, height, err := term.GetSize(fd)
	if err != nil {
		return 0, 0, err
	}
	return width, height, nil
}


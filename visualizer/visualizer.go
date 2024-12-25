package visualizer

import (
	"fmt"
	"os"
	"golang.org/x/term"
)

// RenderFrequencies renders the frequency bars on the terminal and centers the content
func RenderFrequencies(frequencies []float64) {
	maxHeight := 25 // Maximum height of the bars
	numBars := 40   // Number of bars

	// Normalize and scale the frequencies
	scaledHeights := make([]int, numBars)
	for i, freq := range frequencies[:numBars] {
		height := int(freq * float64(maxHeight))
		if height > maxHeight {
			height = maxHeight
		}
		scaledHeights[i] = height
	}

	// Get full terminal size, not tmux pane size
	width, height, err := getTerminalSize()
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		return
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
				fmt.Print("█ ") // Print filled bar
			} else {
				fmt.Print("  ") // Print empty space
			}
		}
	}

	// Reset cursor to the top to avoid scrolling (move cursor back to top-left)
	fmt.Print("\033[H")
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


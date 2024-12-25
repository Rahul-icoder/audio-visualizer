package visualizer

import (
	"fmt"
	"os"
	"golang.org/x/term"
)

// RenderFrequencies renders the frequency bars on the terminal and centers the content
func RenderFrequencies(frequencies []float64) {
	maxHeight := 30 // Maximum height of the bars
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

	// Get terminal size
	width, height, err := getTerminalSize()
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		return
	}

	// Calculate the position to center the content horizontally and vertically
	offsetX := (width - numBars*2) / 2 // The width of each bar is 2 (due to spacing)
	offsetY := (height - maxHeight) / 2 // Center the bars vertically

	// Clear the screen and move the cursor to the top-left corner
	fmt.Print("\033[H\033[2J") // Clear the screen

	// Render bars from top to bottom, starting from the calculated offset
	for row := maxHeight; row > 0; row-- {
		// Move the cursor to the vertical offset position
		fmt.Printf("\033[%d;%dH", offsetY+row, offsetX) 

		// Print the bars for the current row
		for _, height := range scaledHeights {
			if height >= row {
				fmt.Print("█ ") // Print filled bar
			} else {
				fmt.Print("  ") // Print empty space
			}
		}
	}

	// Optionally, move the cursor to the bottom after rendering
	fmt.Print("\033[H") // Optional: Reset cursor to the top
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


package main

import (
	"fmt"
	"log"
	"time"

	"example.com/mod/audio"
	"example.com/mod/visualizer"
)

func main() {
	// Initialize PortAudio before capturing any audio
	err := audio.InitializeAudio()
	if err != nil {
		log.Fatalf("Error initializing audio: %v", err)
	}
	defer audio.TerminateAudio() // Ensure PortAudio is terminated when done

	bufferSize := 1024
	err = runAudioVisualizer(bufferSize)
	if err != nil {
		log.Fatalf("Error running audio visualizer: %v", err)
	}
}

func runAudioVisualizer(bufferSize int) error {
	for {
		// Step 1: Capture audio
		buffer, err := audio.CaptureAudio(bufferSize)
		if err != nil {
			return fmt.Errorf("Error capturing audio: %v", err)
		}

		// Step 2: Process audio
		frequencies := audio.ProcessAudio(buffer)

		// Step 3: Render visualizer
		visualizer.RenderFrequencies(frequencies)

		// Step 4: Control frame rate (10 FPS)
		time.Sleep(100 * time.Millisecond)
	}
}


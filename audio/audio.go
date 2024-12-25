package audio

import (
	"fmt"
	"github.com/gordonklaus/portaudio"
	"math/cmplx"
	"github.com/mjibson/go-dsp/fft"
)

// InitializePortAudio initializes the PortAudio library
func InitializeAudio() error {
	if err := portaudio.Initialize(); err != nil {
		return fmt.Errorf("error initializing PortAudio: %v", err)
	}
	return nil
}

// TerminatePortAudio terminates the PortAudio library
func TerminateAudio() {
	portaudio.Terminate()
}

// CaptureAudio captures audio from the default input stream
func CaptureAudio(bufferSize int) ([]float32, error) {
	buffer := make([]float32, bufferSize)
	stream, err := portaudio.OpenDefaultStream(1, 0, 44100, len(buffer), buffer)
	if err != nil {
		return nil, fmt.Errorf("error opening audio stream: %v", err)
	}
	defer stream.Close()

	if err := stream.Start(); err != nil {
		return nil, fmt.Errorf("error starting audio stream: %v", err)
	}
	defer stream.Stop()

	if err := stream.Read(); err != nil {
		return nil, fmt.Errorf("error reading from audio stream: %v", err)
	}

	return buffer, nil
}

// ProcessAudio processes the captured audio buffer to extract frequency magnitudes
func ProcessAudio(buffer []float32) []float64 {
	complexInput := make([]complex128, len(buffer))
	for i, v := range buffer {
		complexInput[i] = complex(float64(v), 0)
	}

	frequencies := fft.FFT(complexInput)
	magnitudes := make([]float64, len(frequencies))
	maxMagnitude := 0.0

	for i, f := range frequencies {
		magnitude := cmplx.Abs(f)
		magnitudes[i] = magnitude
		if magnitude > maxMagnitude {
			maxMagnitude = magnitude
		}
	}

	// Normalize magnitudes to a range of 0 to 1
	for i := range magnitudes {
		magnitudes[i] /= maxMagnitude
	}

	return magnitudes
}


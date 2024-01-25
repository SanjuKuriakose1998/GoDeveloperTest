package main

import (
	"fmt"
	"sync"
)

// declares a slice to represent the shared data buffer
var sharedBuffer []byte

// declares a mutual exclusion lock to synchronize access to the slice
var mutex sync.Mutex

// Defines a struct to represent the configuration for the number of M and N
type Routine struct {
	M int
	N int
}

func main() {

	// allocates and initializes a shared buffer with a size of 20 bytes
	sharedBuffer = make([]byte, 20)

	// creates a slice of ‘Routine’ instances with different configuration
	routines := []Routine{
		{M: 8, N: 2},
		{M: 8, N: 8},
		{M: 8, N: 16},
		{M: 2, N: 8},
	}

	for _, r := range routines {
		runRoutine(r)
	}
}

// defines the ‘runRoutine’ function that launches reader and writer goroutines
func runRoutine(r Routine) {
	for i := 0; i < r.M; i++ {
		go reader(i)
	}

	for j := 0; j < r.N; j++ {
		go writer(j)
	}

	select {} // Used to keep the main goroutine running indefinitely
}

// defines the ‘writer’ function that simulates a writer goroutine
func writer(wIndex int) {
	for {
		mutex.Lock()
		index := wIndex % len(sharedBuffer)
		sharedBuffer[index] = byte(wIndex)
		fmt.Println("written to buffer: ", wIndex)
		mutex.Unlock()
	}
}

// defines the ‘reader’ function that simulates a reader goroutine
func reader(rIndex int) {
	for {
		mutex.Lock()
		index := rIndex % len(sharedBuffer)
		value := sharedBuffer[index]
		fmt.Println("read from buffer: ", value)
		mutex.Unlock()
	}
}

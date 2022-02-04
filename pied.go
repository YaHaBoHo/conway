package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, gridChunk []byte) {
	defer wg.Done()
	for i := 0; i < len(gridChunk); i++ {
		gridChunk[i]++
		time.Sleep(10 * time.Millisecond)
	}
}

func pied() {
	// Initialize data
	grid := make([]byte, 1000)
	// Initialize threading
	dmz := 3
	threads := 8
	var wg sync.WaitGroup
	chunkSize := (len(grid) - threads*dmz) / threads
	if chunkSize < dmz {
		// Playing it safe
		panic("Too many threads for grid size!")
	}
	// First pass
	start := time.Now()
	for row := 0; row < len(grid); row += chunkSize {
		wg.Add(1)
		if row+chunkSize < len(grid) {
			go worker(&wg, grid[row+dmz:row+chunkSize])
		} else {
			go worker(&wg, grid[row+dmz:])
		}
	}
	wg.Wait()
	// Second pass
	for row := 0; row < len(grid); row += chunkSize {
		wg.Add(1)
		worker(&wg, grid[row:row+dmz])
	}
	wg.Wait()
	// Done
	fmt.Println(time.Now().Sub(start).Seconds())
}

func main() {
	pied()
}

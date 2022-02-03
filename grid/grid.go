package grid

import (
	"conway/config"
	"math/rand"
)

func EmptyGrid() *[config.GridSize][config.GridSize]byte {
	var grid [config.GridSize][config.GridSize]byte
	return &grid
}

func PentominoGrid() *[config.GridSize][config.GridSize]byte {
	// Initialize
	grid := EmptyGrid()
	// Insert a pentomino pattern
	gridCenter := config.GridSize / 2
	grid[gridCenter-1][gridCenter+1] = 1
	grid[gridCenter-1][gridCenter] = 1
	grid[gridCenter][gridCenter] = 1
	grid[gridCenter][gridCenter-1] = 1
	grid[gridCenter+1][gridCenter] = 1
	// Done
	return grid
}

func RandomGrid(density float32) *[config.GridSize][config.GridSize]byte {
	// Initialize
	grid := EmptyGrid()
	for row := 0; row < config.GridSize; row++ {
		for col := 0; col < config.GridSize; col++ {
			if rand.Float32() < density {
				grid[row][col] = 1
			}
		}
	}
	return grid
}

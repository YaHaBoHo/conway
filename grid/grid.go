package grid

import (
	"conway/config"
	"math/rand"
)

func EmptyGrid() *[config.GridSize * config.GridSize]byte {
	var grid [config.GridSize * config.GridSize]byte
	return &grid
}

func RandomGrid(density float32) *[config.GridSize * config.GridSize]byte {
	// Initialize
	grid := EmptyGrid()
	for row := 0; row < config.GridSize; row++ {
		for col := 0; col < config.GridSize; col++ {
			if rand.Float32() < density {
				grid[row*config.GridSize+col] = 1
			}
		}
	}
	return grid
}

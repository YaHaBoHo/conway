package main

import (
	"conway/config"
	"conway/grid"
	"fmt"
	"math"
	"time"
)

func updateAlive(grid *[config.GridSize * config.GridSize]byte, row int, col int, center int) byte {
	var numNeighbors byte = 0
	var left, right, top, bottom int
	// Left
	if col == 0 {
		left = config.GridSize - 1
	} else {
		left = -1
	}
	// Right
	if col == config.GridSize-1 {
		right = -config.GridSize + 1
	} else {
		right = 1
	}
	// Top
	if row == 0 {
		top = config.GridLen - config.GridSize
	} else {
		top = -config.GridSize
	}
	// Bottom
	if row == config.GridSize-1 {
		bottom = config.GridSize - config.GridLen
	} else {
		bottom = config.GridSize
	}
	// Go (lol)
	numNeighbors += grid[center+left+top]
	numNeighbors += grid[center+top]
	numNeighbors += grid[center+right+top]
	numNeighbors += grid[center+left]
	numNeighbors += grid[center+right]
	numNeighbors += grid[center+left+bottom]
	numNeighbors += grid[center+bottom]
	if numNeighbors < 1 {
		return 0
	}
	numNeighbors += grid[center+right+bottom]
	if numNeighbors == 2 || numNeighbors == 3 {
		return 1
	}
	return 0
}

func updateDead(grid *[config.GridSize * config.GridSize]byte, row int, col int, center int) byte {
	var numNeighbors byte = 0
	var left, right, top, bottom int
	// Left
	if col == 0 {
		left = config.GridSize - 1
	} else {
		left = -1
	}
	// Right
	if col == config.GridSize - 1 {
		right = - config.GridSize + 1
	} else {
		right = 1
	}
	// Top
	if row == 0 {
		top = config.GridLen - config.GridSize
	} else {
		top = -config.GridSize
	}
	// Bottom
	if row == config.GridSize-1 {
		bottom = config.GridSize - config.GridLen
	} else {
		bottom = config.GridSize
	}
	// Go (lol)
	numNeighbors += grid[center+left+top]
	numNeighbors += grid[center+top]
	numNeighbors += grid[center+right+top]
	numNeighbors += grid[center+left]
	numNeighbors += grid[center+right]
	numNeighbors += grid[center+left+bottom]
	if numNeighbors < 1 {
		return 0
	}
	numNeighbors += grid[center+bottom]
	numNeighbors += grid[center+right+bottom]
	if numNeighbors == 3 {
		return 1
	}
	return 0
}

func main() {
	// Initialize
	fmt.Println("Loading grid...")
	// var startRle = rle.FromFile(config.InputFile)
	// var mainGrid = rle.ToGrid(startRle)
	var mainGrid = grid.RandomGrid(0.3)
	// Run
	fmt.Println("Starting simulation...")
	start := time.Now()
	for round := 0; round < config.Rounds; round++ {
		tempGrid := grid.EmptyGrid()
		for row := 0; row < config.GridSize; row++ {
			for col := 0; col < config.GridSize; col++ {
				var idx = row*config.GridSize + col
				if mainGrid[idx] == 1 {
					tempGrid[idx] = updateAlive(mainGrid, row, col, idx)
				} else {
					tempGrid[idx] = updateDead(mainGrid, row, col, idx)
				}
			}
		}
		mainGrid = tempGrid
	}
	taken := time.Now().Sub(start).Seconds()
	fmt.Println("Simulation completed.")
	//if rle.FromGrid(mainGrid).Data == rle.FromFile(config.CheckFile).Data {
	totalCells := float64(config.Rounds * config.GridSize * config.GridSize)
	fmt.Printf("Size              : %vx%v\n", config.GridSize, config.GridSize)
	fmt.Printf("Time              : %v s\n", math.Round(taken*100)/100)
	fmt.Printf("Rounds            : %v\n", config.Rounds)
	fmt.Printf("Round time (avg)  : %v ms\n", int(1000*taken)/config.Rounds)
	fmt.Printf("Cell rate         : %v Mc/s\n", math.Round(totalCells/(taken*10000))/100)
	//} else {
	//	panic("Pattern differs from expected!")
	//}
}

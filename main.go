package main

import (
	"fmt"
	"time"
)

const GridSize int = 512
const Rounds int = 1000

func startingGrid() [GridSize][GridSize]int {
	// Initialize
	var grid [GridSize][GridSize]int
	// Insert a pentomino pattern
	gridCenter := GridSize / 2
	grid[gridCenter-1][gridCenter-1] = 1
	grid[gridCenter-1][gridCenter] = 1
	grid[gridCenter][gridCenter] = 1
	grid[gridCenter][gridCenter+1] = 1
	grid[gridCenter+1][gridCenter] = 1
	// Done
	return grid
}

func updateAlive(grid *[GridSize][GridSize]int, row int, col int) int {
	var numNeighbors int = 0
	var cMin = (col - 1 + GridSize) % GridSize
	var cMax = (col + 1) % GridSize
	// [1] Top row, no check needed
	var rMin = (row - 1 + GridSize) % GridSize
	numNeighbors += grid[rMin][cMin]
	numNeighbors += grid[rMin][col]
	numNeighbors += grid[rMin][cMax]
	// [2] Center row, ignore center column, check > 3 after each operation
	numNeighbors += grid[row][cMin]
	if numNeighbors > 3 {
		return 0
	}
	numNeighbors += grid[row][cMax]
	if numNeighbors > 3 {
		return 0
	}
	// [3] Bottom row, check > 3 after each operation
	var rMax = (row + 1) % GridSize
	numNeighbors += grid[rMax][cMin]
	if numNeighbors > 3 {
		return 0
	}
	numNeighbors += grid[rMax][col]
	if numNeighbors > 3 {
		return 0
	}
	numNeighbors += grid[rMax][cMax]
	if numNeighbors > 3 {
		return 0
	}
	// [4] Final > 1 check
	if numNeighbors > 1 {
		return 1
	}
	return 0
}

func updateDead(grid *[GridSize][GridSize]int, row int, col int) int {
	var numNeighbors int = 0
	var cMin = (col - 1 + GridSize) % GridSize
	var cMax = (col + 1) % GridSize
	// [1] Top row, no check needed
	var rMin = (row - 1 + GridSize) % GridSize
	numNeighbors += grid[rMin][cMin]
	numNeighbors += grid[rMin][col]
	numNeighbors += grid[rMin][cMax]
	// [2] Bottom row, check > 3 after each operation
	var rMax = (row + 1) % GridSize
	numNeighbors += grid[rMax][cMin]
	if numNeighbors > 3 {
		return 0
	}
	numNeighbors += grid[rMax][col]
	if numNeighbors > 3 {
		return 0
	}
	numNeighbors += grid[rMax][cMax]
	if numNeighbors > 3 {
		return 0
	}
	// [3] Center row, ignore center column, check > 3 after first operation, stop before if 0 neighbours
	if numNeighbors < 1 {
		return 0
	}
	numNeighbors += grid[row][cMin]
	if numNeighbors > 3 {
		return 0
	}
	numNeighbors += grid[row][cMax]
	// [4] Final == 3 check
	if numNeighbors == 3 {
		return 1
	}
	return 0
}

func main() {
	// Initialize
	var grid = startingGrid()
	// Run
	start := time.Now()
	for round := 0; round < Rounds; round++ {
		var newGrid [GridSize][GridSize]int
		for row := 0; row < GridSize; row++ {
			for col := 0; col < GridSize; col++ {
				if grid[row][col] == 1 {
					newGrid[row][col] = updateAlive(&grid, row, col)
				} else {
					newGrid[row][col] = updateDead(&grid, row, col)
				}
			}
		}
		grid = newGrid
	}
	taken := time.Now().Sub(start).Seconds()
	fmt.Println("Time : ", taken)
	fmt.Println("RPS  : ", float64(Rounds)/taken)
	fmt.Println("CPS  : ", float64(Rounds*GridSize*GridSize)/taken)
}

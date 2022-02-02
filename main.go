package main

import (
	"conway/config"
	"fmt"
	"math"
	"time"
)

func emptyGrid() *[config.GridSize][config.GridSize]int {
	var grid [config.GridSize][config.GridSize]int
	return &grid
}

func startingGrid() *[config.GridSize][config.GridSize]int {
	// Initialize
	var grid [config.GridSize][config.GridSize]int
	// Insert a pentomino pattern
	gridCenter := config.GridSize / 2
	grid[gridCenter-1][gridCenter+1] = 1
	grid[gridCenter-1][gridCenter] = 1
	grid[gridCenter][gridCenter] = 1
	grid[gridCenter][gridCenter-1] = 1
	grid[gridCenter+1][gridCenter] = 1
	// Done
	return &grid
}

func updateAlive(grid *[config.GridSize][config.GridSize]int, row int, col int) int {
	var numNeighbors int = 0
	var cMin = (col - 1 + config.GridSize) % config.GridSize
	var cMax = (col + 1) % config.GridSize
	// [1] Top row, no check needed
	var rMin = (row - 1 + config.GridSize) % config.GridSize
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
	var rMax = (row + 1) % config.GridSize
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

func updateDead(grid *[config.GridSize][config.GridSize]int, row int, col int) int {
	var numNeighbors int = 0
	var cMin = (col - 1 + config.GridSize) % config.GridSize
	var cMax = (col + 1) % config.GridSize
	// [1] Top row, no check needed
	var rMin = (row - 1 + config.GridSize) % config.GridSize
	numNeighbors += grid[rMin][cMin]
	numNeighbors += grid[rMin][col]
	numNeighbors += grid[rMin][cMax]
	// [2] Bottom row, check > 3 after each operation
	var rMax = (row + 1) % config.GridSize
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
	for round := 0; round < config.Rounds; round++ {
		newGrid := emptyGrid()
		for row := 0; row < config.GridSize; row++ {
			for col := 0; col < config.GridSize; col++ {
				if grid[row][col] == 1 {
					newGrid[row][col] = updateAlive(grid, row, col)
				} else {
					newGrid[row][col] = updateDead(grid, row, col)
				}
			}
		}
		grid = newGrid
	}
	taken := time.Now().Sub(start).Seconds()
	totalCells := float64(config.Rounds * config.GridSize * config.GridSize)
	fmt.Printf("Size              : %vx%v\n", config.GridSize, config.GridSize)
	fmt.Printf("Time              : %v s\n", math.Round(taken*100)/100)
	fmt.Printf("Rounds            : %v\n", config.Rounds)
	fmt.Printf("Round time (avg)  : %v ms\n", 1000*int(taken)/config.Rounds)
	fmt.Printf("Cell rate         : %v Mc/s\n", math.Round(totalCells/(taken*10000))/100)
	// fmt.Println(utilities.ToRLE(grid))
}

package main

import (
	"conway/config"
	"conway/utilities"
	"fmt"
	"math"
	"time"
)

func updateAlive(grid *[config.GridSize][config.GridSize]byte, row int, col int) byte {
	var numNeighbors byte = 0
	var col0 = (col - 1 + config.GridSize) % config.GridSize
	var col2 = (col + 1) % config.GridSize
	var row0 = (row - 1 + config.GridSize) % config.GridSize
	numNeighbors += grid[row0][col0]
	numNeighbors += grid[row0][col]
	numNeighbors += grid[row0][col2]
	numNeighbors += grid[row][col0]
	numNeighbors += grid[row][col2]
	var row2 = (row + 1) % config.GridSize
	numNeighbors += grid[row2][col0]

	numNeighbors += grid[row2][col]
	if numNeighbors < 1 {
		return 0
	}
	numNeighbors += grid[row2][col2]
	if numNeighbors > 1 && numNeighbors < 4 {
		return 1
	}
	return 0
}

func updateDead(grid *[config.GridSize][config.GridSize]byte, row int, col int) byte {
	var numNeighbors byte = 0
	var col0 = (col - 1 + config.GridSize) % config.GridSize
	var col2 = (col + 1) % config.GridSize
	var row0 = (row - 1 + config.GridSize) % config.GridSize
	numNeighbors += grid[row0][col0]
	numNeighbors += grid[row0][col]
	numNeighbors += grid[row0][col2]
	numNeighbors += grid[row][col0]
	numNeighbors += grid[row][col2]
	var row2 = (row + 1) % config.GridSize
	numNeighbors += grid[row2][col0]
	if numNeighbors < 1 {
		return 0
	}
	numNeighbors += grid[row2][col]
	numNeighbors += grid[row2][col2]
	if numNeighbors == 3 {
		return 1
	}
	return 0
}

func main() {
	// Initialize
	var grid = utilities.RandomGrid(0.3)
	// Run
	start := time.Now()
	for round := 0; round < config.Rounds; round++ {
		newGrid := utilities.EmptyGrid()
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
	fmt.Printf("Round time (avg)  : %v ms\n", int(1000*taken)/config.Rounds)
	fmt.Printf("Cell rate         : %v Mc/s\n", math.Round(totalCells/(taken*10000))/100)
	// fmt.Println(utilities.ToRLE(grid))
}

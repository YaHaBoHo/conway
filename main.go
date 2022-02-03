package main

import (
	"conway/config"
	"conway/grid"
	"conway/rle"
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
	fmt.Println("Loading grid...")
	var startRle = rle.FromFile(config.InputFile)
	var mainGrid = rle.ToGrid(startRle)
	// Run
	fmt.Println("Starting simulation...")
	start := time.Now()
	for round := 0; round < config.Rounds; round++ {
		tempGrid := grid.EmptyGrid()
		for row := 0; row < config.GridSize; row++ {
			for col := 0; col < config.GridSize; col++ {
				if mainGrid[row][col] == 1 {
					tempGrid[row][col] = updateAlive(mainGrid, row, col)
				} else {
					tempGrid[row][col] = updateDead(mainGrid, row, col)
				}
			}
		}
		mainGrid = tempGrid
	}
	taken := time.Now().Sub(start).Seconds()
	fmt.Println("Simulation completed.")
	if rle.FromGrid(mainGrid).Data == rle.FromFile(config.CheckFile).Data {
		totalCells := float64(config.Rounds * config.GridSize * config.GridSize)
		fmt.Printf("Size              : %vx%v\n", config.GridSize, config.GridSize)
		fmt.Printf("Time              : %v s\n", math.Round(taken*100)/100)
		fmt.Printf("Rounds            : %v\n", config.Rounds)
		fmt.Printf("Round time (avg)  : %v ms\n", int(1000*taken)/config.Rounds)
		fmt.Printf("Cell rate         : %v Mc/s\n", math.Round(totalCells/(taken*10000))/100)
	} else {
		panic("Pattern differs from expected!")
	}
}

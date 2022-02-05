package conway

import (
	"sync"
)

const DmzRows int = 3

func cellSurroundings(world *World, row int, col int) (int, int, int, int) {
	// Init
	var left, right, top, bottom int
	// Left
	if col == 0 {
		left = world.Size - 1
	} else {
		left = -1
	}
	// Right
	if col == world.Size-1 {
		right = -world.Size + 1
	} else {
		right = 1
	}
	// Top
	if row == 0 {
		top = world.Cells - world.Size
	} else {
		top = -world.Size
	}
	// Bottom
	if row == world.Cells-world.Size {
		bottom = world.Size - world.Cells
	} else {
		bottom = world.Size
	}
	return left, right, top, bottom
}

func SetStagedCell(world *World, row int, col int) {
	// Prepare
	center := row + col
	left, right, top, bottom := cellSurroundings(world, row, col)
	// Cell
	world.Next[center] += 0x1
	// Neighbors
	world.Next[center+left+top] += 0x2
	world.Next[center+top] += 0x2
	world.Next[center+right+top] += 0x2
	world.Next[center+left] += 0x2
	world.Next[center+right] += 0x2
	world.Next[center+left+bottom] += 0x2
	world.Next[center+bottom] += 0x2
	world.Next[center+right+bottom] += 0x2
}

func updateStagedCell(world *World, row int, col int) {
	cell := world.Grid[row+col]
	// Process cells
	if cell == 6 {
		SetStagedCell(world, row, col)
	} else if cell == 5 || cell == 7 {
		SetStagedCell(world, row, col)
	}

}

func updateChunk(wg *sync.WaitGroup, world *World, rowStart int, rowEnd int) {
	defer wg.Done()
	for row := rowStart; row < rowEnd; row += world.Size {
		for col := 0; col < world.Size; col++ {
			updateStagedCell(world, row, col)
		}
	}

}

func updateGridConcurrently(world *World, concurrency int) {
	// Prepare concurrency
	var wg sync.WaitGroup
	dmzCells := world.Size * DmzRows
	chunkSize := (world.Cells - dmzCells*concurrency) / concurrency
	// chunkSize := (len(grid) - threads*dmz) / threads
	// world.Size * ((world.Size - concurrency*DmzRows) / concurrency)
	if chunkSize < dmzCells {
		// Playing it safe
		panic("Too many threads for grid size!")
	}
	// Stage world
	Stage(world)
	// Compute : First pass
	for row := 0; row < world.Cells; row += chunkSize {
		wg.Add(1)
		if row+chunkSize < world.Cells {
			go updateChunk(&wg, world, row+dmzCells, row+chunkSize)
		} else {
			go updateChunk(&wg, world, row+dmzCells, world.Cells)
		}
	}
	wg.Wait()
	// Compute : Second pass
	for row := 0; row < world.Cells; row += chunkSize {
		wg.Add(1)
		go updateChunk(&wg, world, row, row+dmzCells)
	}
	wg.Wait()
	// Done
	Commit(world)
}

func updateGrid(world *World) {
	// Stage
	Stage(world)
	// Compute
	for row := 0; row < world.Cells; row += world.Size {
		for col := 0; col < world.Size; col++ {
			updateStagedCell(world, row, col)
		}
	}
	// Done
	Commit(world)
}

func Simulate(world *World, rounds int, concurrency int) {
	if concurrency > 1 {
		for r := 0; r < rounds; r++ {
			updateGridConcurrently(world, concurrency)
		}
	} else {
		for r := 0; r < rounds; r++ {
			updateGrid(world)
		}
	}
}

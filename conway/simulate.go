package conway

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
	neighbors := cell >> 1
	// Process cells
	if cell&1 == 0 {
		// Dead cells
		if neighbors == 3 {
			SetStagedCell(world, row, col)
		}
	} else {
		// Living cells
		if neighbors == 2 || neighbors == 3 {
			SetStagedCell(world, row, col)
		}
	}
}

func updateGrid(world *World) {
	Stage(world)
	for row := 0; row < world.Cells; row += world.Size {
		for col := 0; col < world.Size; col++ {
			updateStagedCell(world, row, col)
		}
	}
	// Commit
	Commit(world)
}

func Simulate(world *World, rounds int) {
	for r := 0; r < rounds; r++ {
		updateGrid(world)
		// fmt.Println(r, Density(world))
	}
}

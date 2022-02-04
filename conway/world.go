package conway

import (
	"fmt"
	"math/rand"
)

type World struct {
	Grid  []byte
	Next  []byte
	Cells int
	Size  int
}

func Stage(world *World) {
	world.Next = EmptyGrid(world.Cells)
}

func Commit(world *World) {
	world.Grid = world.Next
}

func EmptyGrid(cells int) []byte {
	return make([]byte, cells)
}

func EmptyWorld(size int) *World {
	worldCells := size * size
	return &World{
		Grid:  EmptyGrid(worldCells),
		Next:  EmptyGrid(worldCells),
		Cells: worldCells,
		Size:  size}
}

func RandomWorld(size int, density float32) *World {
	// Initialize
	world := EmptyWorld(size)
	// Generate
	for row := 0; row < world.Cells; row += world.Size {
		for col := 0; col < world.Size; col++ {
			if rand.Float32() < density {
				SetStagedCell(world, row, col)
			}
		}
	}
	// Commit
	Commit(world)
	// Done
	return world
}

func Density(world *World) float64 {
	var cellCount float64
	for i := 0; i < world.Size; i++ {
		cellCount += float64(world.Grid[i] & 1)
	}
	return cellCount / float64(world.Size)
}

func PrintState(world *World) {
	// Separator
	separator := "o"
	for col := 0; col < world.Size; col++ {
		separator += "-"
	}
	separator += "o"
	// World
	fmt.Println(separator)
	for row := 0; row < world.Cells; row += world.Size {
		textLine := "|"
		for col := 0; col < world.Size; col++ {
			//textLine += fmt.Sprintf("%d", world.Grid[col+row]>>1)
			if world.Grid[col+row]&1 == 1 {
				textLine += "+"
			} else {
				textLine += " "
			}
		}
		fmt.Println(textLine + "|")
	}
	fmt.Println(separator)
}

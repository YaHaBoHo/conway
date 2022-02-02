package utilities

import (
	"conway/config"
	"fmt"
)

type RLE struct {
	header string
	data   string
}

func rleItem(state byte, count int) string {
	// Get tag
	var rleTag string
	if state == 1 {
		rleTag = "o"
	} else {
		rleTag = "b"
	}
	// If count > 1, return <count><tag>
	if count > 1 {
		return fmt.Sprintf("%d%s", count, rleTag)
	}
	// Else, just return <tag>
	return rleTag
}

func rleEnding(count int) string {
	// If count > 1, return <count>$
	if count > 1 {
		return fmt.Sprintf("%d$", count)
	}
	// Else, just return $
	return "$"
}

func ToRLE(grid *[config.GridSize][config.GridSize]byte) RLE {
	var rleData string
	var rleHeader = fmt.Sprintf("x=%d,y=%d,rule=B3/S23", config.GridSize, config.GridSize)
	var newLines = 0
	for row := 0; row < config.GridSize; row++ {
		// Initialize line
		var rleRow string
		// Initialize accumulators
		var aState = grid[row][0]
		var aCount = 0
		for col := 0; col < config.GridSize; col++ {
			if aState == grid[row][col] {
				aCount += 1
			} else {
				rleRow += rleItem(aState, aCount)
				aState = grid[row][col]
				aCount = 1
			}
		}
		// Dump leftover accumulator
		if aCount > 0 && aState == 1 {
			rleRow += rleItem(aState, aCount)
		}
		// If line has cells, append to rle
		if rleRow == "" {
			newLines += 1
		} else {
			// Dump line endings
			if newLines > 0 {
				rleData += rleEnding(newLines)
			}
			newLines = 1
			rleData += rleRow
		}
	}
	// Finalize RLE
	rleData += "!"
	// Done
	return RLE{header: rleHeader, data: rleData}
}

func FromRLE(rle string) *[config.GridSize][config.GridSize]int {
	var grid [config.GridSize][config.GridSize]int
	// Some stuff ...
	return &grid
}

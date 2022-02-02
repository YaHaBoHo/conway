package utilities

import (
	"conway/config"
	"strconv"
)

func toRleItem(state int, count int) string {
	// Get tag
	var rleTag string
	if state == 1 {
		rleTag = "o"
	} else {
		rleTag = "b"
	}
	// If count > 1, return <count><tag>
	if count > 1 {
		return strconv.Itoa(count) + rleTag
	}
	// Else, just return <tag>
	return rleTag
}

func toRleEnding(count int) string {
	// If count > 1, return <count>$
	if count > 1 {
		return strconv.Itoa(count) + "$"
	}
	// Else, just return $
	return "$"
}

func ToRLE(grid *[config.GridSize][config.GridSize]int) string {
	var rle string
	var newLines = 0
	for row := 0; row < config.GridSize; row++ {
		// Initialize line
		var rleLine string
		// Initialize accumulators
		var aState = grid[row][0]
		var aCount = 0
		for col := 0; col < config.GridSize; col++ {
			if aState == grid[row][col] {
				aCount += 1
			} else {
				rleLine += toRleItem(aState, aCount)
				aState = grid[row][col]
				aCount = 1
			}
		}
		// Dump leftover accumulator
		if aCount > 0 && aState == 1 {
			rleLine += toRleItem(aState, aCount)
		}
		// If line has cells, append to rle
		if rleLine == "" {
			newLines += 1
		} else {
			// Dump line endings
			if newLines > 0 {
				rle += toRleEnding(newLines)
			}
			newLines = 1
			rle += rleLine
		}
	}
	// Finalize RLE
	rle += "!"
	// Done
	return rle
}

func FromRLE(rle string) *[config.GridSize][config.GridSize]int {
	var grid [config.GridSize][config.GridSize]int
	// Some stuff ...
	return &grid
}

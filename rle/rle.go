package rle

import (
	"bufio"
	"conway/conway"
	"conway/utilities"
	"fmt"
	"os"
	"strings"
)

const FileLineSize = 70
const SummaryLength = 128

const LiveChar = "o"
const DeadChar = "b"
const NewLine = "$"
const End = "!"

type RLE struct {
	Header string
	Data   string
}

func Summary(rle RLE) string {
	rleLen := len(rle.Data)
	if rleLen > SummaryLength {
		return fmt.Sprintf("%s...%s (%d)", rle.Data[:SummaryLength/2], rle.Data[rleLen-SummaryLength/2:], rleLen)
	}
	return fmt.Sprintf("%s (%d)", rle.Data, rleLen)
}

func formatData(rawText string) string {
	// Format RLE
	var fmtText string
	for c := 0; c < len(rawText); c += FileLineSize {
		var cMax = c + FileLineSize
		if cMax < len(rawText) {
			fmtText += rawText[c:c+FileLineSize] + "\n"
		} else {
			fmtText += rawText[c:]
		}
	}
	return fmtText
}

func rleItem(state byte, count int) string {
	// Get tag
	var rleTag string
	if state == 1 {
		rleTag = LiveChar
	} else {
		rleTag = DeadChar
	}
	// If count > 1, return <count><tag>
	if count > 1 {
		return fmt.Sprintf("%d%s", count, rleTag)
	}
	// Else, just return <tag>
	return rleTag
}

func rleNewLine(count int) string {
	// If count > 1, return <count>$
	if count > 1 {
		return fmt.Sprintf("%d%s", count, NewLine)
	}
	// Else, just return $
	return NewLine
}

func FromWorld(world *conway.World) RLE {
	var rleData string
	var rleHeader = fmt.Sprintf("x=%d,y=%d,rule=B3/S23", world.Size, world.Size)
	var newLines = 0
	for row := 0; row < world.Cells; row += world.Size {
		// Initialize line
		var rleRow string
		// Initialize buffers
		var bufferState = world.Grid[row] & 1
		var bufferCount = 0
		for col := 0; col < world.Size; col++ {
			cellState := world.Grid[row+col] & 1
			if bufferState == cellState {
				bufferCount += 1
			} else {
				rleRow += rleItem(bufferState, bufferCount)
				bufferState = cellState
				bufferCount = 1
			}
		}
		// Dump leftover buffer
		if bufferCount > 0 && bufferState == 1 {
			rleRow += rleItem(bufferState, bufferCount)
		}
		// If line has cells, append to rle
		if rleRow == "" {
			newLines += 1
		} else {
			// Dump line endings
			if newLines > 0 {
				rleData += rleNewLine(newLines)
			}
			newLines = 1
			rleData += rleRow
		}
	}
	// Finalize RLE
	rleData += End
	// Done
	return RLE{Header: rleHeader, Data: rleData}
}

func ToWorld(rle RLE) *conway.World {
	// Header
	header := strings.Split(rle.Header, ",")
	if len(header) < 3 {
		panic("RLE header is too short")
	}
	headerCols := strings.Split(header[0], "=")
	headerRows := strings.Split(header[1], "=")
	if len(headerCols) < 2 || len(headerRows) < 2 {
		panic("Size descriptor in RLE header is invalid.")
	}
	rleRows := utilities.ParseInt(headerRows[1])
	rleCols := utilities.ParseInt(headerCols[1])
	if rleRows != rleCols {
		panic(fmt.Sprintf("Only square grids are supported yet, not %dx%d!", rleRows, rleCols))
	}
	// Data
	var col = 0
	var row = 0
	var numBuffer string
	world := conway.EmptyWorld(rleRows)
	for _, c := range rle.Data {
		char := string(c)
		if char == End {
			break
		}
		if char == LiveChar || char == DeadChar || char == NewLine {
			// Buffer?
			var repeat = 1
			if len(numBuffer) > 0 {
				repeat = utilities.ParseInt(numBuffer)
			}
			if char == LiveChar {
				// Living cell(s)
				for i := 0; i < repeat; i++ {
					conway.SetStagedCell(world, row, col+i)
				}
				col += repeat
			} else if char == DeadChar {
				// Dead cell(s)
				col += repeat
			} else if char == NewLine {
				// New line
				col = 0
				row += repeat * world.Size
			}
			numBuffer = ""
		} else if char == " " {
			// Ignore whitespaces
		} else {
			numBuffer += char
		}
	}
	// Commit
	conway.Commit(world)
	// Done
	return world
}

func FromFile(filePath string) RLE {
	// Open file
	fileData, err := os.ReadFile(filePath)
	utilities.CheckError(err)
	// Process data
	lines := strings.Split(string(fileData), "\n")
	if len(lines) < 2 {
		panic(fmt.Sprintf("Not enough lines in RLE file %s", filePath))
	}
	// Extract data
	var rleData string
	for _, line := range lines[1:] {
		rleData += line
	}
	// Done
	return RLE{Header: lines[0], Data: rleData}
}

func ToFile(rle RLE, filePath string) {
	// Open file
	file, err := os.Create(filePath)
	utilities.CheckError(err)
	// Prepare closure
	defer func(file *os.File) {
		utilities.CheckError(file.Close())
	}(file)
	// Prepare writer
	writer := bufio.NewWriter(file)
	// Write header
	_, err = writer.WriteString(rle.Header + "\n")
	utilities.CheckError(err)
	// Write data
	_, err = writer.WriteString(formatData(rle.Data))
	utilities.CheckError(err)
	// Done
	utilities.CheckError(writer.Flush())
	return
}

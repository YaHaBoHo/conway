package rle

import (
	"bufio"
	"conway/config"
	"conway/utilities"
	"fmt"
	"os"
	"strings"
)

const MaxLength = 70
const LiveChar = "o"
const DeadChar = "b"
const NewLine = "$"
const End = "!"

type RLE struct {
	Header string
	Data   string
}

func formatData(rawText string) string {
	// Format RLE
	var fmtText string
	for c := 0; c < len(rawText); c += MaxLength {
		var cMax = c + MaxLength
		if cMax < len(rawText) {
			fmtText += rawText[c:c+MaxLength] + "\n"
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

func FromGrid(grid *[config.GridSize][config.GridSize]byte) RLE {
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

func ToGrid(rle RLE) *[config.GridSize][config.GridSize]byte {
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
	rleCols := utilities.ParseInt(headerCols[1])
	rleRows := utilities.ParseInt(headerRows[1])
	if (rleRows != config.GridSize) || (rleCols != config.GridSize) {
		panic(fmt.Sprintf("RLE size (%dx%d) does not match grid size (%dx%d)!", rleRows, rleCols, config.GridSize, config.GridSize))
	}
	// Data
	var grid [config.GridSize][config.GridSize]byte
	var col int
	var row int
	var numBuffer string
	for _, c := range rle.Data {
		char := string(c)
		if char == End {
			break
		}
		if char == LiveChar || char == DeadChar || char == NewLine {
			// Buffer?
			var repeat int = 1
			if len(numBuffer) > 0 {
				repeat = utilities.ParseInt(numBuffer)
			}
			if char == LiveChar {
				// Dead cell(s)
				for i := 0; i < repeat; i++ {
					grid[row][col+i] = 1
				}
				col += repeat
			} else if char == DeadChar {
				// Dead cell(s)
				col += repeat
			} else if char == NewLine {
				// New line
				col = 0
				row += repeat
			}
			numBuffer = ""
		} else if char == " " {
			// Ignore whitespaces
		} else {
			numBuffer += char
		}
	}
	// Done
	return &grid
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

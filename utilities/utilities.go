package utilities

import "strconv"

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func ParseInt(numString string) int {
	parsedInt64, err := strconv.ParseInt(numString, 10, 64)
	CheckError(err)
	return int(parsedInt64)
}

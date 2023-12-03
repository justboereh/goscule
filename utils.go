package goscule

import (
	"math"
	"unicode"
)

func includes[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}

	return false
}

func splitByCase(str string, seperators []string) []string {
	if len(seperators) < 1 {
		seperators = STR_SPLITTER[:]
	}

	parts := []string{}

	buff := ""
	previousUpper := false
	previousSplitter := false

	for _, char := range str {
		isSplitter := includes(seperators, string(char))

		if isSplitter {
			parts = append(parts, buff)
			buff = ""
			previousSplitter = false

			continue
		}

		isUpper := unicode.IsUpper(char)

		if !previousSplitter && !previousUpper && isUpper {
			parts = append(parts, buff)
			buff = string(char)
			previousUpper = isUpper
			continue
		}

		if !previousSplitter && previousUpper && !isUpper && len(buff) > 1 {
			parts = append(parts, buff[0:int(math.Max(0, float64(len(buff)-1)))])
			buff = buff[:len(buff)-1] + string(char)
			previousUpper = isUpper
			continue
		}

		buff += string(char)
		previousUpper = isUpper
		previousSplitter = isSplitter
	}

	parts = append(parts, buff)

	return parts

}

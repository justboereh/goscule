package goscule

import (
	"math"
	"unicode"
	"unicode/utf8"
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

	for _, char := range []rune(str) {
		isSplitter := includes(seperators, string(char))

		if isSplitter {
			parts = append(parts, buff)
			buff = ""
			previousSplitter = false

			continue
		}

		isUpper := unicode.IsUpper(char)

		if previousSplitter == false && previousUpper == false && isUpper == true {
			parts = append(parts, buff)
			buff = string(char)
			previousUpper = isUpper
			continue
		}

		if previousSplitter == false && previousUpper == true && isUpper == false && len(buff) > 1 {
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

func upperFirst(str string) string {
	if len(str) < 1 {
		return ""
	}

	_, i := utf8.DecodeRuneInString(str)
	therest := str[i:]
	first := unicode.ToUpper([]rune(str)[0])

	return string(first) + therest
}

func lowerFirst(str string) string {
	if len(str) < 1 {
		return ""
	}

	_, i := utf8.DecodeRuneInString(str)
	therest := str[i:]
	first := unicode.ToLower([]rune(str)[0])

	return string(first) + therest
}

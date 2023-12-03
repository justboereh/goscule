package goscule

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

var STR_SPLITTER = []string{"-", "_", "/", "."}

func Case(arr []string, fn func(str string) string, joiner string) string {
	result := ""

	for i, str := range arr {
		result += fn(str)

		if i < len(arr) {
			result += joiner
		}
	}

	return result
}

func Pascal(str string) string {
	if len(str) < 1 {
		return ""
	}

	return PascalArray(splitByCase(str, STR_SPLITTER))
}

func PascalArray(arr []string) string {
	if len(arr) < 1 {
		return ""
	}

	return Case(arr, func(str string) string {
		return UpperFirst(str)
	}, "")
}

func Camel(str string) string {
	if len(str) < 1 {
		return ""
	}

	return LowerFirst(Pascal(str))
}

func CamelArray(arr []string) string {
	if len(arr) < 1 {
		return ""
	}

	return LowerFirst(PascalArray(arr))
}

func Kebab(str string) string {
	if len(str) < 1 {
		return ""
	}

	return KebabArray(splitByCase(str, STR_SPLITTER))
}

func KebabArray(arr []string) string {
	if len(arr) < 1 {
		return ""
	}

	return Case(arr, func(str string) string {
		return strings.ToLower(str)
	}, "-")
}

func Snake(str string) string {
	if len(str) < 1 {
		return ""
	}

	return SnakeArray(splitByCase(str, STR_SPLITTER))
}

func SnakeArray(arr []string) string {
	if len(arr) < 1 {
		return ""
	}

	return Case(arr, func(str string) string {
		return strings.ToLower(str)
	}, "_")
}

func UpperFirst(str string) string {
	if len(str) < 1 {
		return ""
	}

	_, i := utf8.DecodeRuneInString(str)
	therest := str[i:]
	first := unicode.ToUpper([]rune(str)[0])

	return string(first) + therest
}

func LowerFirst(str string) string {
	if len(str) < 1 {
		return ""
	}

	_, i := utf8.DecodeRuneInString(str)
	therest := str[i:]
	first := unicode.ToLower([]rune(str)[0])

	return string(first) + therest
}

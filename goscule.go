package goscule

var STR_SPLITTER = [4]string{"-", "_", "/", "."}

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

}

func PascalArray(arr []string) string {
	if len(arr) < 1 {
		return ""
	}

	return Case(arr, func(str string) string {
		return upperFirst(str)
	}, "")
}

func Camel(str string) string {
	if len(str) < 1 {
		return ""
	}

}

func Kebab(str string) string {
	if len(str) < 1 {
		return ""
	}

}

func Snake(str string) string {
	if len(str) < 1 {
		return ""
	}

}

func Upper(str string) string {
	if len(str) < 1 {
		return ""
	}

}

func Lower(str string) string {
	if len(str) < 1 {
		return ""
	}

}

package uniq

import (
	// "fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type FlagsStruct struct {
	C    bool
	D    bool
	U    bool
	F    int
	S    int
	I    bool
	Help bool
}

func Uniq(input []string, flags FlagsStruct) (output []string) {
	var result []string
	var prevStr string
	count := 0
	strMap := map[string]int{}

	for i, str := range input {
		newStr := str

		if flags.I {
			newStr = strings.ToLower(str)
		}

		if flags.F > 0 {
			arr := strings.SplitN(newStr, " ", flags.F+1)
			newStr = arr[len(arr)-1]
		}

		if flags.S > 0 && utf8.RuneCountInString(newStr) > 0 {
			arr := strings.SplitN(newStr, "", flags.S+1)
			newStr = arr[len(arr)-1]
		}

		if _, ok := strMap[newStr]; !ok {
			result = append(result, str)
		}

		count++
		if strings.Compare(newStr, prevStr) != 0 && i != 0 {
			if tmpStr, add := addToReslut(prevStr, count, flags); add {
				output = append(output, tmpStr)
			}
			count = 0
		}
		prevStr = newStr
	}
	count++
	if tmpStr, add := addToReslut(prevStr, count, flags); add {
		output = append(output, tmpStr)
	}

	return
}

func addToReslut(str string, count int, flags FlagsStruct) (string, bool) {
	if flags.C {
		return (strconv.Itoa(count) + " " + str), true
	} else if flags.D {
		if count > 1 {
			return str, true
		}
	} else if flags.U {
		if count == 1 {
			return str, true
		}
	} else {
		return str, true
	}
	return "", false
}

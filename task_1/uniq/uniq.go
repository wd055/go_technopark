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
	var prevStr, tmpStr string
	count := 0

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

		if i == 0 {
			tmpStr = str
		}

		if strings.Compare(newStr, prevStr) != 0 && i != 0 {
			if tmpStr, add := addToReslut(tmpStr, count, flags); add {
				output = append(output, tmpStr)
			}
			tmpStr = str
			count = 1
		} else {
			count++
		}
		prevStr = newStr
	}

	if tmpStr, add := addToReslut(tmpStr, count, flags); add {
		output = append(output, tmpStr)
	}

	return
}

func addToReslut(str string, count int, flags FlagsStruct) (string, bool) {
	switch {
	case flags.C:
		return (strconv.Itoa(count) + " " + str), true
	case flags.D && count > 1:
		return str, true
	case flags.U && count == 1:
		return str, true
	}
	return str, true
}

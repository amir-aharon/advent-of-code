package aoc

import (
	"strconv"
	"strings"
	"unicode"
)

func FetchSliceOfIntsInString(line string) []int {
	nums := []int{}
	var build strings.Builder
	isNegative := false

	for _, char := range line {
		if unicode.IsDigit(char) {
			build.WriteRune(char)
		}
		if char == '-' {
			isNegative = true
		}
		if (char == ' ' || char == ',' || char == '~' || char == '|') && build.Len() != 0 {
			localNum, err := strconv.ParseInt(build.String(), 10, 64)
			if err != nil {
				panic(err)
			}
			if isNegative {
				localNum *= -1
			}
			nums = append(nums, int(localNum))
			build.Reset()
			isNegative = false
		}
	}
	if build.Len() != 0 {
		localNum, err := strconv.ParseInt(build.String(), 10, 64)
		if err != nil {
			panic(err)
		}
		if isNegative {
			localNum *= -1
		}
		nums = append(nums, int(localNum))
		build.Reset()
	}
	return nums
}

func SplitStringSliceByEmptyLines(input []string) [][]string {
    res := [][]string{}

    currentPart := []string{}
    for _, line := range input {
        if strings.TrimSpace(line) == "" {
            if len(currentPart) > 0 {
                res = append(res, currentPart)
            }
            currentPart = []string{}
        } else {
            currentPart = append(currentPart, line)
        }
    }
    if len(currentPart) > 0 {
        res = append(res, currentPart)
    }
    return res
}

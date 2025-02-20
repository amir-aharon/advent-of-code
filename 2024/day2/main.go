package main

import (
	"fmt"

	aoc "aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")
	//intArrays := getNumArrayFromColumns(input)

	fmt.Println("answer to part 1: ", getSafeReportsCount(input))
	fmt.Println("answer to part 2: ", getSafeReportsCountWithDeletion(input))

}

func isSafeReport(report []int) bool {
	flagDecrease, flagIncrease := false, false

	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]

		if diff > 0 {
			flagIncrease = true
		} else if diff < 0 {
			flagDecrease = true
		} else {
			return false
		}

		if flagDecrease && flagIncrease {
			return false
		}

		if diff < -3 || 3 < diff {
			return false
		}
	}
	return true
}

func getCopyWithDeletedIndex(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}
	newSlice := make([]int, 0, len(slice)-1)
	newSlice = append(newSlice, slice[:index]...)
	newSlice = append(newSlice, slice[index+1:]...)
    //fmt.Println(index, slice, newSlice)
	return newSlice
}

func isSafeReportWithDeletion(report []int) bool {
	if isSafeReport(report) {
		return true
	}
	for i := 0; i < len(report); i++ {
		if isSafeReport(getCopyWithDeletedIndex(report, i)) {
			return true
		}
	}
	return false
}

func getSafeReportsCount(reports []string) int {
	count := 0
	for _, report := range reports {
		if isSafeReport(aoc.FetchSliceOfIntsInString(report)) {
			count += 1
		}
	}
	return count
}

func getSafeReportsCountWithDeletion(reports []string) int {
	count := 0
	for _, report := range reports {
		if isSafeReportWithDeletion(aoc.FetchSliceOfIntsInString(report)) {
			count += 1
		}
	}
	return count
}

package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	aoc "aoc/library"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")

	inputParts := aoc.SplitStringSliceByEmptyLines(input)
	if len(inputParts) != 2 {
		panic("The input is invalid")
	}

	ors, ptps := []OrderingRule{}, []PageToProduce{}
	for _, or := range inputParts[0] {
		ors = append(ors, *NewOrderingRule(or))
	}
	for _, ptp := range inputParts[1] {
		ptps = append(ptps, NewPageToProduce(ptp))
	}

	fmt.Println("answer to part 1: ", Part1(ptps, ors))
	fmt.Println("answer to part 2: ", Part2(ptps, ors))

}

func Part1(ptps []PageToProduce, ors []OrderingRule) int {
	onlyValidPtps := FilterPtps(ptps, ors, true)
	return SumMiddlePageNums(onlyValidPtps)
}

func Part2(ptps []PageToProduce, ors []OrderingRule) int {
	onlyInvalidParts := FilterPtps(ptps, ors, false)
	sorted := []PageToProduce{}
	for _, ptp := range onlyInvalidParts {
		sorted = append(sorted, RuleSort(ptp, ors))
	}
	return SumMiddlePageNums(sorted)
}

func RuleSort(ptp PageToProduce, ors []OrderingRule) PageToProduce {
    for !IsValid(ptp, ors) {
        for _, or := range ors {
            if !or.Validate(ptp) {
                preLoc := aoc.FindIndexInSlice(ptp, or.PreNum)
                postLoc := aoc.FindIndexInSlice(ptp, or.PostNum)
                ptp[preLoc] = or.PostNum
                ptp[postLoc] = or.PreNum
            }
        }
    }
    return ptp
}

func IsValid(ptp PageToProduce, ors []OrderingRule) bool {
	for _, or := range ors {
		if !or.Validate(ptp) {
			return false
		}
	}
	return true
}

func SumMiddlePageNums(input []PageToProduce) int {
	sum := 0
	for _, ptp := range input {
		sum += ptp[len(ptp)/2]
	}
	return sum
}

func FilterPtps(ptps []PageToProduce, ors []OrderingRule, validOnes bool) []PageToProduce {
	res := []PageToProduce{}
	for _, ptp := range ptps {
        isValid := IsValid(ptp, ors)
        if (validOnes && isValid) || (!validOnes && !isValid) {
            res = append(res, ptp)
        }
	}
	return res
}

type OrderingRule struct {
	PreNum  int
	PostNum int
}

func NewOrderingRule(textualRule string) *OrderingRule {
	rule := OrderingRule{}
	reader := strings.NewReader(strings.TrimSpace(textualRule))
	_, err := fmt.Fscanf(reader, "%d|%d", &rule.PreNum, &rule.PostNum)
	if err != nil {
		panic(err)
	}
	return &rule
}

func (rule *OrderingRule) Validate(ptp PageToProduce) bool {
	preLoc := aoc.FindIndexInSlice(ptp, rule.PreNum)
	postLoc := aoc.FindIndexInSlice(ptp, rule.PostNum)
	if postLoc == -1 || preLoc == -1 {
		return true
	}
	if postLoc < preLoc {
		return false
	}
	return true
}

func NewPageToProduce(textuelPtp string) PageToProduce {
	strParts := regexp.MustCompile(`\d+`).FindAllString(textuelPtp, -1)
	res := []int{}
	for _, part := range strParts {
		intPart, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}
		res = append(res, intPart)
	}
	return res
}

type PageToProduce []int

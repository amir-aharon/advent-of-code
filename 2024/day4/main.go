package main

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	aoc "aoc/library"
)

const (
	XMAS = "XMAS"
)

func main() {
	input := aoc.ReadFileLineByLine("input.txt")

    fmt.Println("answer 1: ", CountXMAS(input))
	fmt.Println("answer 2: ", CountXMASSquares(input))

}

func CountXMAS(input []string) int {
	if len(input) == 0 {
		return 0
	}
	count := 0
	spans := getSpans(input)
	for _, span := range spans {
		count += countInSpan(span)
	}
	return count
}

func getSpans(input []string) []string {
	height := len(input)
	width := len(input[0])
	columns := []string{}
	rows := []string{}
	diagonals := []string{}
	reverseDiagonals := []string{}

	// columns
	for col := 0; col < width; col++ {
		var s strings.Builder
		for row := 0; row < height; row++ {
			s.WriteByte(input[row][col])
		}
		columns = append(columns, s.String())
		s.Reset()
	}

	// rows
	for row := 0; row < height; row++ {
		var s strings.Builder
		for col := 0; col < width; col++ {
			s.WriteByte(input[row][col])
		}
		rows = append(rows, s.String())
		s.Reset()
	}

	// diagonals
	for col := 0; col < width; col++ {
		var s strings.Builder
		for relRow, relCol := 0, col; relRow < height && relCol < width; relRow, relCol = relRow+1, relCol+1 {
			s.WriteByte(input[relRow][relCol])
		}
		diagonals = append(diagonals, s.String())
		s.Reset()
	}

	for row := 1; row < height; row++ {
		var s strings.Builder
		for relRow, relCol := row, 0; relRow < height && relCol < width; relRow, relCol = relRow+1, relCol+1 {
			s.WriteByte(input[relRow][relCol])
		}
		diagonals = append(diagonals, s.String())
		s.Reset()
	}

	// reverse diagonals
	for col := width - 1; col >= 0; col-- {
		var s strings.Builder
		for relRow, relCol := 0, col; relRow < height && relCol >= 0; relRow, relCol = relRow+1, relCol-1 {
			s.WriteByte(input[relRow][relCol])
		}
		reverseDiagonals = append(reverseDiagonals, s.String())
		s.Reset()
	}

	for row := 1; row < height; row++ {
		var s strings.Builder
		for relRow, relCol := row, width-1; relRow < height && relCol >= 0; relRow, relCol = relRow+1, relCol-1 {
			s.WriteByte(input[relRow][relCol])
		}
		reverseDiagonals = append(reverseDiagonals, s.String())
		s.Reset()
	}
	return slices.Concat(columns, rows, diagonals, reverseDiagonals)
}

func countInSpan(span string) int {
	count := 0
	re := regexp.MustCompile(XMAS)
	count += len(re.FindAllString(span, -1))
	count += len(re.FindAllString(reverse(span), -1))
	return count
}

func reverse(str string) string {
	rns := []rune(str)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func CountXMASSquares(matrix []string) int {
    subSquares := getSquareSize3(matrix)
    possibilities := [][]string{}
    for i, square := range subSquares {
        fmt.Println(i, len(subSquares))
        possibilities = slices.Concat(possibilities, getAllSquareRotations(square))
    }
    fmt.Println(possibilities)
    count := 0
    for _, possibility := range possibilities {
        if isXmasSquare(possibility) {
            count += 1
        }
    }
    return count
}

func isXmasSquare(square []string) bool {
    fmt.Println(square)
	XMAS_SQUARE := []string{
		"M.S", ".A.", "M.S",
	}
	for row := 0; row < len(square); row++ {
		for col := 0; col < len(square[0]); col++ {
			if XMAS_SQUARE[row][col] != '.' && XMAS_SQUARE[row][col] != square[row][col] {
				return false
			}
		}
	}
	return true
}

func getSquareSize3(matrix []string) [][]string {
	height := len(matrix)
	width := len(matrix[0])
	res := [][]string{}
	for row := 0; row < height-2; row++ {
		for col := 0; col < width-2; col++ {
            sq := []string {
                string([]byte{matrix[row][col],matrix[row][col+1],matrix[row][col+2],}),
                string([]byte{matrix[row+1][col],matrix[row+1][col+1],matrix[row+1][col+2],}),
                string([]byte{matrix[row+2][col],matrix[row+2][col+1],matrix[row+2][col+2],}),
            }
            res = append(res, sq)
		}
	}
	return res
}

func getAllSquareRotations(matrix []string) [][]string {
	res := [][]string{}
	deg0 := matrix
	res = append(res, deg0)
	deg180 := []string{}
	for _, row := range deg0 {
		deg180 = append(deg180, reverse(row))
	}
	res = append(res, deg180)

	deg90 := []string{}
	for col := 0; col < len(matrix[0]); col++ {
		var s strings.Builder
		for row := 0; row < len(matrix); row++ {
			s.WriteByte(matrix[row][col])
		}
		deg90 = append(deg90, s.String())
		s.Reset()
	}
	res = append(res, deg90)

	deg270 := []string{}
	for _, row := range deg90 {
		deg270 = append(deg270, reverse(row))
	}
	res = append(res, deg270)

	return res
}

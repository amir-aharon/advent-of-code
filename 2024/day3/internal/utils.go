package internal

import (
	"bufio"
	"log"
	"os"
)

func CombineRegex(tokens []TokenDefinition) string {
	expr := ""
	for _, token := range tokens {
		expr += token.Pattern
		expr += "|"
	}
	return expr[:len(expr)-1]
}

func ReadFileLineByLine(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var output []string
	scanner := bufio.NewScanner(file)
	const maxCapacity = 512 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return output
}

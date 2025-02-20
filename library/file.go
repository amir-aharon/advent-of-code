package aoc

import (
	"bufio"
	"log"
	"os"
)

func ReadFileLineByLine(path string) []string {
	// correct way to open a file
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var output []string

	// create a scanner for file
	scanner := bufio.NewScanner(file)

	// handle files up to 512 KB
	const maxCapacity = 512 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	// scan line by line
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	// how scanner inform errors
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return output
}

func WriteToFile(path string, content string) {
    // create file
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
    defer f.Close()

    // write content
    if _, err = f.WriteString(content); err != nil {
        panic(err)
    }
}

package main

import (
	"fmt"
	"interpreter/internal"
	"strings"
)

func main() {
	input := strings.Join(internal.ReadFileLineByLine("input.txt"), "")
    reg := internal.TokenRegistry
    tokens := internal.Tokenize(input, reg)
    endContext := internal.Evaluate(tokens)
	fmt.Println("answer: ", endContext.Sum)
}

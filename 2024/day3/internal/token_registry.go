package internal

import "interpreter/internal/token"

type TokenDefinition struct {
	Pattern     string
	Constructor func(string) token.Token
}

var TokenRegistry = []TokenDefinition{
	{Pattern: `do\(\)`, Constructor: token.NewDoToken},
	{Pattern: `don't\(\)`, Constructor: token.NewDontToken},
	{Pattern: `mul\(\d{1,3},\d{1,3}\)`, Constructor: token.NewMulToken},
}

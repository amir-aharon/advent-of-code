package token

import (
	"fmt"
	"interpreter/internal/context"
	"strings"
)

type MulToken struct {
	op1 int
	op2 int
}

func (t MulToken) Evaluate(ctx *context.Context) {
	if ctx.Enabled {
		ctx.Sum += t.op1 * t.op2
	}
}

func NewMulToken(val string) Token {
	t := MulToken{}
    reader := strings.NewReader(val)
	_, err := fmt.Fscanf(reader, "mul(%d,%d)", &t.op1, &t.op2)
	if err != nil {
		panic(err)
	}
	return t
}

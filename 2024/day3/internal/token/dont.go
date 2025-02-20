package token

import (
	"interpreter/internal/context"
)

type DontToken struct{}

func NewDontToken(val string) Token {
	return DontToken{}
}

func (t DontToken) Evaluate(ctx *context.Context) {
	ctx.Enabled = false
}

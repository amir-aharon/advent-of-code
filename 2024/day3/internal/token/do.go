package token

import (
	"interpreter/internal/context"
)

type DoToken struct{}

func NewDoToken(val string) Token {
	return DoToken{}
}

func (t DoToken) Evaluate(ctx *context.Context) {
	ctx.Enabled = true
}

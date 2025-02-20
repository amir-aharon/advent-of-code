package token

import (
	"interpreter/internal/context"
)

type Token interface {
	Evaluate(ctx *context.Context)
}

package internal

import (
	"interpreter/internal/context"
	"interpreter/internal/token"
	"regexp"
)

func Tokenize(text string, possibleTokens []TokenDefinition) []token.Token {
    re := regexp.MustCompile(CombineRegex(possibleTokens))
    textualTokens := re.FindAllString(text, -1)
    tokens := []token.Token{}

    for _, textualToken := range textualTokens{
        for _, possibleToken := range possibleTokens {
            if regexp.MustCompile(possibleToken.Pattern).MatchString(textualToken) {
                token := possibleToken.Constructor(textualToken)
                tokens = append(tokens, token)
                break
            }
        }
    }
    return tokens
}

func Evaluate(tokens []token.Token) context.Context {
    ctx := context.Context{Enabled: true, Sum: 0}
    for _, token := range tokens {
        token.Evaluate(&ctx)
    }
    return ctx
}

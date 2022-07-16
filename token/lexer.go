package token

import "fmt"

type Token struct {
	Type    string
	Literal string
}

var lexedTokens map[string]Token

func (t Token) String() string {
	return fmt.Sprintf("%s %s", t.Type, t.Literal)
}

func (t Token) AddToken(tokenType string, token string) {
	lexedTokens[token] = Token{tokenType, token}
}

func (t Token) GetToken(token string) Token {
	return lexedTokens[token]
}

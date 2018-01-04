package token

type TokenType string

type Token struct {
	Type    TokenType // The type of the current token
	Literal string    // The literal value of the current token
}

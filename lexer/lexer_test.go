package lexer

import (
	"github.com/bootun/tun/token"
	"testing"
)

func TestNextToken(t *testing.T) {

	src := `+=(){}`

	var tests = []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.PLUS, "+"},
		{token.ASSIGN, "="},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
	}

	lexer := New(src)
	for i, tt := range tests {
		tok := lexer.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. excepted: %q, got: %q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. excepted: %q, got: %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

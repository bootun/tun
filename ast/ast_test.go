package ast

import (
	"testing"

	"github.com/bootun/tun/token"
)

func TestString(t *testing.T) {
	program := &Program{Statements: []Statement{
		&LetStatement{
			Token: token.Token{
				Literal: "let",
				Type:    token.LET,
			},
			Name: &Identifier{
				Token: token.Token{
					Literal: "myVar",
					Type:    token.IDENT,
				},
				Value: "myVar",
			},
			Value: &Identifier{
				Token: token.Token{
					Literal: "anotherVar",
					Type:    token.IDENT,
				},
				Value: "anotherVar",
			},
		},
	}}
	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}

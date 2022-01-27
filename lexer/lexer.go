package lexer

import "github.com/bootun/tun/token"

type Lexer struct {
	input        string
	position     int // current position in input
	readPosition int
	ch           byte
}

func (l *Lexer) NextToken() token.Token {
	tok := token.Token{
		Literal: "",
		Type:    "",
	}
	// TODO: implement me
	return tok
}

func (l *Lexer) readChar() {
	// TODO: implement me
}

func New(src string) *Lexer {
	// TODO: fix position
	l := &Lexer{input: src}
	return l
}

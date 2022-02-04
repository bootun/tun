package lexer

import "github.com/bootun/tun/token"

type Lexer struct {
	input        string
	curLine      int
	curColumn    int  // FIXME: incorrect
	position     int  // current position in input(point current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok.Type = token.EQ
			tok.Literal = string(ch) + string(l.ch)

		} else {
			tok = newToken(token.ASSIGN, l.ch, l.curLine, l.curColumn)
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok.Type = token.NOT_EQ
			tok.Literal = string(ch) + string(l.ch)
			tok.Line = l.curLine
			tok.Column = l.curColumn
		} else {
			tok = newToken(token.BANG, l.ch, l.curLine, l.curColumn)
		}
	case '-':
		tok = newToken(token.MINUS, l.ch, l.curLine, l.curColumn)
	case '/':
		tok = newToken(token.SLASH, l.ch, l.curLine, l.curColumn)
	case '*':
		tok = newToken(token.ASTERISK, l.ch, l.curLine, l.curColumn)
	case '>':
		tok = newToken(token.GT, l.ch, l.curLine, l.curColumn)
	case '<':
		tok = newToken(token.LT, l.ch, l.curLine, l.curColumn)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch, l.curLine, l.curColumn)
	case '(':
		tok = newToken(token.LPAREN, l.ch, l.curLine, l.curColumn)
	case ')':
		tok = newToken(token.RPAREN, l.ch, l.curLine, l.curColumn)
	case ',':
		tok = newToken(token.COMMA, l.ch, l.curLine, l.curColumn)
	case '+':
		tok = newToken(token.PLUS, l.ch, l.curLine, l.curColumn)
	case '{':
		tok = newToken(token.LBRACE, l.ch, l.curLine, l.curColumn)
	case '}':
		tok = newToken(token.RBRACE, l.ch, l.curLine, l.curColumn)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
		tok.Line = l.curLine
		tok.Column = l.curColumn
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookUpIdent(tok.Literal)
			tok.Line = l.curLine
			tok.Column = l.curColumn
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			tok.Line = l.curLine
			tok.Column = l.curColumn
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch, l.curLine, l.curColumn)
			return tok
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func isLetter(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_'
}

// readIdentifier read character until is not letter
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func newToken(typ token.Type, lit byte, line, column int) token.Token {
	return token.Token{Type: typ, Literal: string(lit), Line: line, Column: column}
}

// readChar make lexer pointer advance
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
		return
	}
	l.ch = l.input[l.readPosition]
	l.position = l.readPosition
	l.readPosition++
	l.curColumn++
}

// readNumber read character until is not number
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// New create a Lexer by src, the lexer has been initialized
func New(src string) *Lexer {
	l := &Lexer{input: src, curColumn: 1, curLine: 1}
	l.readChar()
	return l
}

// skipWhitespace just eat up the whitespace and newline character
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		if l.ch == '\n' {
			l.curLine++
			l.curColumn = 1
		}
		l.readChar()
	}
}

// isDigit return true if ch is digit
func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

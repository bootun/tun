package lexer

import "github.com/bootun/tun/token"

type Lexer struct {
	// TODO: migrate to io.Reader
	input        string
	curLine      int
	curColumn    int
	position     int // current position in input(point current char)
	readPosition int // current reading position in input (after current char)
	// TODO: migrate to rune
	ch byte // current char under examination
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()
	column := l.curColumn
	line := l.curLine
	// TODO: reduce the number of newToken function
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok.Type = token.EQ
			tok.Literal = string(ch) + string(l.ch)
			tok.Line = line
			tok.Column = column
		} else {
			tok = newToken(token.ASSIGN, l.ch, line, column)
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
			tok = newToken(token.BANG, l.ch, line, column)
		}
	case '-':
		tok = newToken(token.MINUS, l.ch, line, column)
	case '/':
		tok = newToken(token.SLASH, l.ch, line, column)
	case '*':
		tok = newToken(token.ASTERISK, l.ch, line, column)
	case '>':
		tok = newToken(token.GT, l.ch, line, column)
	case '<':
		tok = newToken(token.LT, l.ch, line, column)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch, line, column)
	case '(':
		tok = newToken(token.LPAREN, l.ch, line, column)
	case ')':
		tok = newToken(token.RPAREN, l.ch, line, column)
	case ',':
		tok = newToken(token.COMMA, l.ch, line, column)
	case '+':
		tok = newToken(token.PLUS, l.ch, line, column)
	case '{':
		tok = newToken(token.LBRACE, l.ch, line, column)
	case '}':
		tok = newToken(token.RBRACE, l.ch, line, column)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
		tok.Line = line
		tok.Column = column
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookUpIdent(tok.Literal)
			tok.Line = line
			tok.Column = column
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			tok.Line = line
			tok.Column = column
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch, line, column)
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
	l := &Lexer{input: src, curColumn: 0, curLine: 1}
	l.readChar()
	return l
}

// skipWhitespace just eat up the whitespace and newline character
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		if l.ch == '\n' {
			l.curLine++
			l.curColumn = 0
		}
		l.readChar()
	}
}

// isDigit return true if ch is digit
func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

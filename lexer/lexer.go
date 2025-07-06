package lexer

import "github.com/mgill25/monkey-go/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // (after current char)
	ch           byte // current char under examination
}

// constructor
func New(input string) *Lexer {
	l := &Lexer{input: input}
	// read one char during lexer construction otherwise we'll fail
	// to get the "next" token - since we haven't read anything yet.
	l.readChar()
	return l
}

// Lexer methods

// readChar: reads the next character into the lexer's ch slot
// and increments the next index by 1
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// Produces a token based on char read in current byte slot
// and then reads a new char into that
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.readChar()
	return tok
}

// helper wrapper function to build tokens
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

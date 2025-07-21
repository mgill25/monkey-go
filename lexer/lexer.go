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

// peekChar() is similar to readChar() except it does not increment
// our position pointers, and neither pushes the read char into l.ch
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

// Produces a token based on char read in current byte slot
// and then reads a new char into that
func (l *Lexer) NextToken() token.Token {
parseToken:
	var tok token.Token
	l.skipWhitespace()
	switch l.ch {
	case '=':
		if string(l.peekChar()) == "=" {
			oldCh := l.ch
			l.readChar()
			literal := string(oldCh) + string(l.ch)
			tok = newTokenWithStr(token.EQ, literal)
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			oldCh := l.ch
			l.readChar()
			literal := string(oldCh) + string(l.ch)
			tok = newTokenWithStr(token.NOT_EQ, literal)
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/': // found first SLASH
		if l.peekChar() == '/' { // Check if the next one is also a SLASH
			l.skipComment()
			goto parseToken
		} else {
			tok = newToken(token.SLASH, l.ch)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) skipComment() {
	for l.ch != '\n' && l.ch != '\r' && l.ch != 0 {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// Monkey lang only supports integers for now
// we are ignoring floats, hexadecimals etc.
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// helper wrapper function to build tokens
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

func newTokenWithStr(tokenType token.TokenType, literal string) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: literal,
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

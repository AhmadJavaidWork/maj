package lexer

import "github.com/ahmadjavaidwork/maj/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	switch l.ch {
	case 0:
		tok = token.Token{Type: token.EOF, Literal: ""}
	case '=':
		tok = token.Token{Type: token.ASSIGN, Literal: "="}
	case '+':
		tok = token.Token{Type: token.PLUS, Literal: "+"}
	case '-':
		tok = token.Token{Type: token.MINUS, Literal: "-"}
	case '!':
		tok = token.Token{Type: token.BANG, Literal: "!"}
	case '/':
		tok = token.Token{Type: token.SLASH, Literal: "/"}
	case ',':
		tok = token.Token{Type: token.COMMA, Literal: ","}
	case ';':
		tok = token.Token{Type: token.SEMICOLON, Literal: ";"}
	case ':':
		tok = token.Token{Type: token.COLON, Literal: ":"}
	case '*':
		tok = token.Token{Type: token.ASTERISK, Literal: "*"}
	default:
		tok = newToken(token.ILLEGAL, l.ch)
	}

	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func newToken(tokentype token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokentype, Literal: string(ch)}
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		l.readChar()
	}
}

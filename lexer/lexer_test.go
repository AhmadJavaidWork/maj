package lexer

import (
	"testing"

	"github.com/ahmadjavaidwork/maj/token"
)

func TestNextToken(t *testing.T) {
	input := `
	let a = 10;
	let b = 20;
	let c = fn(a, b) {
		return a + b;
	}
	a == b;
	a < b;
	a > b;
	a <= b;
	a >= b;
	let d = "hello, world";
	`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "a"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "b"},
		{token.ASSIGN, "="},
		{token.INT, "20"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "c"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "a"},
		{token.COMMA, ","},
		{token.IDENT, "b"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.IDENT, "a"},
		{token.PLUS, "+"},
		{token.IDENT, "b"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.IDENT, "a"},
		{token.EQ, "=="},
		{token.IDENT, "b"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "a"},
		{token.LT, "<"},
		{token.IDENT, "b"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "a"},
		{token.GT, ">"},
		{token.IDENT, "b"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "a"},
		{token.LT_EQ, "<="},
		{token.IDENT, "b"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "a"},
		{token.GT_EQ, ">="},
		{token.IDENT, "b"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "d"},
		{token.ASSIGN, "="},
		{token.STRING, "hello, world"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - token literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

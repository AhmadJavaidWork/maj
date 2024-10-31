package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	SLASH    = "/"
	ASTERISK = "*"
	LT       = "<"
	GT       = ">"

	EQ     = "=="
	NOT_EQ = "!="
	LT_EQ  = "<="
	GT_EQ  = ">="

	// Delimeters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Identifiers
	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	RETURN   = "RETURN"
	IF       = "IF"
	ELSE     = "ELSE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	} else {
		return IDENT
	}
}

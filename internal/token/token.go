package token

type TokenType string

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	IDENT TokenType = "IDENT"
	INT   TokenType = "INT"

	ASSIGN       TokenType = "="
	PLUS         TokenType = "+"
	MINUS        TokenType = "-"
	BANG         TokenType = "!"
	ASTERISK     TokenType = "*"
	SLASH        TokenType = "/"
	LESS_THAN    TokenType = "<"
	GREATER_THAN TokenType = ">"

	EQUALS            TokenType = "=="
	NOT_EQUALS        TokenType = "!="
	LESS_OR_EQUALS    TokenType = "<="
	GREATER_OR_EQUALS TokenType = ">="

	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"

	LPAREN TokenType = "("
	RPAREN TokenType = ")"
	LBRACE TokenType = "{"
	RBRACE TokenType = "}"

	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
	RETURN   TokenType = "RETURN"
	TRUE     TokenType = "TRUE"
	FALSE    TokenType = "FALSE"
	IF       TokenType = "IF"
	ELSE     TokenType = "ELSE"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
}

func LookupIdent(ident string) TokenType {
	if tokenType, ok := keywords[ident]; ok {
		return tokenType
	}

	return IDENT
}

type Token struct {
	Type    TokenType
	Literal string
}

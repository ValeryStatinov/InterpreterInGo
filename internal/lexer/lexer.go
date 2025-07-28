package lexer

import "github.com/ValeryStatinov/InterpreterInGo/internal/token"

type Lexer struct {
	input        string
	position     int  // current pos in input (points to cur char)
	readPosition int  // current reading position (after current char)
	ch           byte // current char under examination
}

func NewLexer(input string) *Lexer {
	lex := &Lexer{
		input: input,
	}

	lex.readChar()

	return lex
}

func newTokenFromLiteral(tokenType token.TokenType, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal}
}

func newTokenFromChar(tokenType token.TokenType, ch byte) token.Token {
	return newTokenFromLiteral(tokenType, string(ch))
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (lex *Lexer) readChar() {
	if lex.readPosition >= len(lex.input) {
		lex.ch = 0
	} else {
		lex.ch = lex.input[lex.readPosition]
	}

	lex.position = lex.readPosition
	lex.readPosition += 1
}

func (lex *Lexer) peekChar() byte {
	if lex.readPosition >= len(lex.input) {
		return 0
	}

	return lex.input[lex.readPosition]
}

func (lex *Lexer) makeTwoCharToken(tokenType token.TokenType) token.Token {
	ch := lex.ch
	lex.readChar()
	literal := string(ch) + string(lex.ch)

	return newTokenFromLiteral(tokenType, literal)
}

func (lex *Lexer) readIdentifier() string {
	start := lex.position
	for isLetter(lex.peekChar()) || isDigit(lex.peekChar()) {
		lex.readChar()
	}

	return lex.input[start:lex.readPosition]
}

func (lex *Lexer) readNumber() string {
	start := lex.position
	for isDigit(lex.peekChar()) {
		lex.readChar()
	}

	return lex.input[start:lex.readPosition]
}

func (lex *Lexer) eatWhitespace() {
	for lex.ch == ' ' || lex.ch == '\t' || lex.ch == '\n' || lex.ch == '\r' {
		lex.readChar()
	}
}

func (lex *Lexer) NextToken() token.Token {
	lex.eatWhitespace()
	tok := newTokenFromChar(token.ILLEGAL, lex.ch)

	switch lex.ch {
	case '=':
		if lex.peekChar() == '=' {
			tok = lex.makeTwoCharToken(token.EQUALS)
		} else {
			tok = newTokenFromChar(token.ASSIGN, lex.ch)
		}
	case '+':
		tok = newTokenFromChar(token.PLUS, lex.ch)
	case '-':
		tok = newTokenFromChar(token.MINUS, lex.ch)
	case '!':
		if lex.peekChar() == '=' {
			tok = lex.makeTwoCharToken(token.NOT_EQUALS)
		} else {
			tok = newTokenFromChar(token.BANG, lex.ch)
		}
	case '*':
		tok = newTokenFromChar(token.ASTERISK, lex.ch)
	case '/':
		tok = newTokenFromChar(token.SLASH, lex.ch)
	case '<':
		if lex.peekChar() == '=' {
			tok = lex.makeTwoCharToken(token.LESS_OR_EQUALS)
		} else {
			tok = newTokenFromChar(token.LESS_THAN, lex.ch)
		}
	case '>':
		if lex.peekChar() == '=' {
			tok = lex.makeTwoCharToken(token.GREATER_OR_EQUALS)
		} else {
			tok = newTokenFromChar(token.GREATER_THAN, lex.ch)
		}
	case ',':
		tok = newTokenFromChar(token.COMMA, lex.ch)
	case ';':
		tok = newTokenFromChar(token.SEMICOLON, lex.ch)
	case '(':
		tok = newTokenFromChar(token.LPAREN, lex.ch)
	case ')':
		tok = newTokenFromChar(token.RPAREN, lex.ch)
	case '{':
		tok = newTokenFromChar(token.LBRACE, lex.ch)
	case '}':
		tok = newTokenFromChar(token.RBRACE, lex.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(lex.ch) {
			literal := lex.readIdentifier()
			tokenType := token.LookupIdent(literal)

			tok = newTokenFromLiteral(tokenType, literal)
		}

		if isDigit(lex.ch) {
			literal := lex.readNumber()
			tokenType := token.INT

			if isLetter(lex.peekChar()) {
				lex.readChar()
				literal += lex.readIdentifier()
				tokenType = token.ILLEGAL
			}

			tok = newTokenFromLiteral(tokenType, literal)
		}
	}

	lex.readChar()

	return tok
}

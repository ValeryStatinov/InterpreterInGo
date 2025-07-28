package lexer

import (
	"fmt"
	"testing"

	"github.com/ValeryStatinov/InterpreterInGo/internal/token"
)

func TestLexer(t *testing.T) {
	t.Run("NextToken single token", func(t *testing.T) {
		type testcase struct {
			input         string
			expectedToken token.Token
		}

		testcases := []testcase{
			{
				input: "=",
				expectedToken: token.Token{
					Type:    token.ASSIGN,
					Literal: "=",
				},
			},
			{
				input: "+",
				expectedToken: token.Token{
					Type:    token.PLUS,
					Literal: "+",
				},
			},
			{
				input: "-",
				expectedToken: token.Token{
					Type:    token.MINUS,
					Literal: "-",
				},
			},
			{
				input: "!",
				expectedToken: token.Token{
					Type:    token.BANG,
					Literal: "!",
				},
			},
			{
				input: "*",
				expectedToken: token.Token{
					Type:    token.ASTERISK,
					Literal: "*",
				},
			},
			{
				input: "/",
				expectedToken: token.Token{
					Type:    token.SLASH,
					Literal: "/",
				},
			},
			{
				input: "<",
				expectedToken: token.Token{
					Type:    token.LESS_THAN,
					Literal: "<",
				},
			},
			{
				input: ">",
				expectedToken: token.Token{
					Type:    token.GREATER_THAN,
					Literal: ">",
				},
			},
			{
				input: "==",
				expectedToken: token.Token{
					Type:    token.EQUALS,
					Literal: "==",
				},
			},
			{
				input: "!=",
				expectedToken: token.Token{
					Type:    token.NOT_EQUALS,
					Literal: "!=",
				},
			},
			{
				input: "<=",
				expectedToken: token.Token{
					Type:    token.LESS_OR_EQUALS,
					Literal: "<=",
				},
			},
			{
				input: ">=",
				expectedToken: token.Token{
					Type:    token.GREATER_OR_EQUALS,
					Literal: ">=",
				},
			},
			{
				input: ",",
				expectedToken: token.Token{
					Type:    token.COMMA,
					Literal: ",",
				},
			},
			{
				input: ";",
				expectedToken: token.Token{
					Type:    token.SEMICOLON,
					Literal: ";",
				},
			},
			{
				input: "(",
				expectedToken: token.Token{
					Type:    token.LPAREN,
					Literal: "(",
				},
			},
			{
				input: ")",
				expectedToken: token.Token{
					Type:    token.RPAREN,
					Literal: ")",
				},
			},
			{
				input: "{",
				expectedToken: token.Token{
					Type:    token.LBRACE,
					Literal: "{",
				},
			},
			{
				input: "}",
				expectedToken: token.Token{
					Type:    token.RBRACE,
					Literal: "}",
				},
			},
			{
				input: "",
				expectedToken: token.Token{
					Type:    token.EOF,
					Literal: "",
				},
			},
			{
				input: "fn",
				expectedToken: token.Token{
					Type:    token.FUNCTION,
					Literal: "fn",
				},
			},
			{
				input: "let",
				expectedToken: token.Token{
					Type:    token.LET,
					Literal: "let",
				},
			},
			{
				input: "return",
				expectedToken: token.Token{
					Type:    token.RETURN,
					Literal: "return",
				},
			},
			{
				input: "true",
				expectedToken: token.Token{
					Type:    token.TRUE,
					Literal: "true",
				},
			},
			{
				input: "false",
				expectedToken: token.Token{
					Type:    token.FALSE,
					Literal: "false",
				},
			},
			{
				input: "if",
				expectedToken: token.Token{
					Type:    token.IF,
					Literal: "if",
				},
			},
			{
				input: "else",
				expectedToken: token.Token{
					Type:    token.ELSE,
					Literal: "else",
				},
			},
			{
				input: "foobar",
				expectedToken: token.Token{
					Type:    token.IDENT,
					Literal: "foobar",
				},
			},
			{
				input: "12345",
				expectedToken: token.Token{
					Type:    token.INT,
					Literal: "12345",
				},
			},
			{
				input: "abc123ef",
				expectedToken: token.Token{
					Type:    token.IDENT,
					Literal: "abc123ef",
				},
			},
			{
				input: "123abc",
				expectedToken: token.Token{
					Type:    token.ILLEGAL,
					Literal: "123abc",
				},
			},
		}

		for _, tc := range testcases {
			lex := NewLexer(tc.input)
			if tc.expectedToken.Type == token.FUNCTION {
				fmt.Println("here")
			}
			tok := lex.NextToken()
			if tok.Type != tc.expectedToken.Type {
				t.Errorf("expected token type %q, got %q", tc.expectedToken.Type, tok.Type)
			}
			if tok.Literal != tc.expectedToken.Literal {
				t.Errorf("expected token literal %q, got %q", tc.expectedToken.Literal, tok.Literal)
			}
		}
	})

	t.Run("NextToken multiple tokens", func(t *testing.T) {
		type testcase struct {
			input          string
			expectedTokens []token.Token
		}

		testcases := []testcase{
			{
				input: "let five = 5;",
				expectedTokens: []token.Token{
					{Type: token.LET, Literal: "let"},
					{Type: token.IDENT, Literal: "five"},
					{Type: token.ASSIGN, Literal: "="},
					{Type: token.INT, Literal: "5"},
					{Type: token.SEMICOLON, Literal: ";"},
					{Type: token.EOF, Literal: ""},
				},
			},
			{
				input: "let fn legalName(x) {}; let 1illegalName{}",
				expectedTokens: []token.Token{
					{Type: token.LET, Literal: "let"},
					{Type: token.FUNCTION, Literal: "fn"},
					{Type: token.IDENT, Literal: "legalName"},
					{Type: token.LPAREN, Literal: "("},
					{Type: token.IDENT, Literal: "x"},
					{Type: token.RPAREN, Literal: ")"},
					{Type: token.LBRACE, Literal: "{"},
					{Type: token.RBRACE, Literal: "}"},
					{Type: token.SEMICOLON, Literal: ";"},
					{Type: token.LET, Literal: "let"},
					{Type: token.ILLEGAL, Literal: "1illegalName"},
					{Type: token.LBRACE, Literal: "{"},
					{Type: token.RBRACE, Literal: "}"},
					{Type: token.EOF, Literal: ""},
				},
			},
			{
				input: "/ let! if<> else * true; false != == <= >=",
				expectedTokens: []token.Token{
					{Type: token.SLASH, Literal: "/"},
					{Type: token.LET, Literal: "let"},
					{Type: token.BANG, Literal: "!"},
					{Type: token.IF, Literal: "if"},
					{Type: token.LESS_THAN, Literal: "<"},
					{Type: token.GREATER_THAN, Literal: ">"},
					{Type: token.ELSE, Literal: "else"},
					{Type: token.ASTERISK, Literal: "*"},
					{Type: token.TRUE, Literal: "true"},
					{Type: token.SEMICOLON, Literal: ";"},
					{Type: token.FALSE, Literal: "false"},
					{Type: token.NOT_EQUALS, Literal: "!="},
					{Type: token.EQUALS, Literal: "=="},
					{Type: token.LESS_OR_EQUALS, Literal: "<="},
					{Type: token.GREATER_OR_EQUALS, Literal: ">="},
					{Type: token.EOF, Literal: ""},
				},
			},
			{
				input: `let five = 5;
				let ten = 10;
				let add = fn(x, y) {
					return x + y;
				};
				let result = add(five, ten);
				`,
				expectedTokens: []token.Token{
					{Type: token.LET, Literal: "let"},
					{Type: token.IDENT, Literal: "five"},
					{Type: token.ASSIGN, Literal: "="},
					{Type: token.INT, Literal: "5"},
					{Type: token.SEMICOLON, Literal: ";"},
					{Type: token.LET, Literal: "let"},
					{Type: token.IDENT, Literal: "ten"},
					{Type: token.ASSIGN, Literal: "="},
					{Type: token.INT, Literal: "10"},
					{Type: token.SEMICOLON, Literal: ";"},
					{Type: token.LET, Literal: "let"},
					{Type: token.IDENT, Literal: "add"},
					{Type: token.ASSIGN, Literal: "="},
					{Type: token.FUNCTION, Literal: "fn"},
					{Type: token.LPAREN, Literal: "("},
					{Type: token.IDENT, Literal: "x"},
					{Type: token.COMMA, Literal: ","},
					{Type: token.IDENT, Literal: "y"},
					{Type: token.RPAREN, Literal: ")"},
					{Type: token.LBRACE, Literal: "{"},
					{Type: token.RETURN, Literal: "return"},
					{Type: token.IDENT, Literal: "x"},
					{Type: token.PLUS, Literal: "+"},
					{Type: token.IDENT, Literal: "y"},
					{Type: token.SEMICOLON, Literal: ";"},
					{Type: token.RBRACE, Literal: "}"},
					{Type: token.SEMICOLON, Literal: ";"},
					{Type: token.LET, Literal: "let"},
					{Type: token.IDENT, Literal: "result"},
					{Type: token.ASSIGN, Literal: "="},
					{Type: token.IDENT, Literal: "add"},
					{Type: token.LPAREN, Literal: "("},
					{Type: token.IDENT, Literal: "five"},
					{Type: token.COMMA, Literal: ","},
					{Type: token.IDENT, Literal: "ten"},
					{Type: token.RPAREN, Literal: ")"},
					{Type: token.SEMICOLON, Literal: ";"},
					{Type: token.EOF, Literal: ""},
				},
			},
		}

		for i, tc := range testcases {
			lex := NewLexer(tc.input)
			tokens := []token.Token{}

			var tok token.Token
			for tok.Type != token.EOF {
				tok = lex.NextToken()
				tokens = append(tokens, tok)
			}

			if len(tokens) != len(tc.expectedTokens) {
				t.Fatalf("case %d: expected %d tokens, got %d", i, len(tc.expectedTokens), len(tokens))
			}

			for j, expectedToken := range tc.expectedTokens {
				if tokens[j].Type != expectedToken.Type {
					t.Fatalf("case %d: expected token type %q, got %q", i, expectedToken.Type, tokens[j].Type)
				}
				if tokens[j].Literal != expectedToken.Literal {
					t.Fatalf("case %d: expected token literal %q, got %q", i, expectedToken.Literal, tokens[j].Literal)
				}
			}
		}
	})
}

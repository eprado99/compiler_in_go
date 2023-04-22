package regexLexer

import "testing"

// func TestTokens(t *testing.T) {

// 	t.Run("With extra keywords and patterns", func(t *testing.T) {
// 		test := `package main
// 				 i := 0`
// 		got := GetTokens(test, "package", "inferredAssignment0:=")
// 		var want = []Token{
// 			Token{value: "package", tokenType: TokenType("package")},
// 			Token{value: "main", tokenType: Identifier},
// 			Token{value: "i", tokenType: Identifier},
// 			Token{value: ":=", tokenType: TokenType("inferredAssignment")},
// 			Token{value: "0", tokenType: Number},
// 		}

// 	})
// }

// Refactor ideas: Create a class Lexer that will hold patterns and keywords so that I can initialize this and attach GetToken method to this class

func TestToken(t *testing.T) {

	tokenTests := []struct {
		name      string
		lexeme    string
		wantToken Token
	}{
		{name: "package keyword", lexeme: "package", wantToken: Token{value: "package", tokenType: Keyword}},
		{name: "main identifier", lexeme: "main", wantToken: Token{value: "main", tokenType: Identifier}},
		{name: "i identifier", lexeme: "i", wantToken: Token{value: "i", tokenType: Identifier}},
		{name: "= assignment", lexeme: "=", wantToken: Token{value: "=", tokenType: Assignment}},
		{name: "0 number", lexeme: "0", wantToken: Token{value: "0", tokenType: Number}},
	}

	for _, tt := range tokenTests {
		t.Run(tt.name, func(t *testing.T) {
			var gotTokenArr []Token

			GetToken(tt.lexeme, &gotTokenArr)
			if gotTokenArr[0].tokenType != tt.wantToken.tokenType {
				t.Errorf("Lexeme: %s Got %s Want %s", tt.lexeme, gotTokenArr[0].tokenType, tt.wantToken.tokenType)
			}
			if gotTokenArr[0].value != tt.wantToken.value {
				t.Errorf("Lexeme: %s Got %s Want %s", tt.lexeme, gotTokenArr[0].value, tt.wantToken.value)
			}
		})
	}
}

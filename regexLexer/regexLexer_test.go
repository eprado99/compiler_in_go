package regexLexer

import "testing"

func TestToken(t *testing.T) {
	l := Lexer{}
	l.Load()
	tokenTests := []struct {
		name      string
		lexeme    string
		wantToken Token
	}{
		{name: "package keyword", lexeme: "package", wantToken: Token{value: "package", tokenType: Keyword}},
		{name: "main identifier", lexeme: "main", wantToken: Token{value: "main", tokenType: TokenType("identifier")}},
		{name: "i identifier", lexeme: "i", wantToken: Token{value: "i", tokenType: TokenType("identifier")}},
		{name: "= assignment", lexeme: "=", wantToken: Token{value: "=", tokenType: TokenType("assignment")}},
		{name: "0 number", lexeme: "0", wantToken: Token{value: "0", tokenType: TokenType("number")}},
	}

	for _, tt := range tokenTests {
		t.Run(tt.name, func(t *testing.T) {
			var gotTokenArr []Token

			GetToken(l, tt.lexeme, &gotTokenArr)
			if gotTokenArr[0].tokenType != tt.wantToken.tokenType {
				t.Errorf("Lexeme: %s Got %s Want %s", tt.lexeme, gotTokenArr[0].tokenType, tt.wantToken.tokenType)
			}
			if gotTokenArr[0].value != tt.wantToken.value {
				t.Errorf("Lexeme: %s Got %s Want %s", tt.lexeme, gotTokenArr[0].value, tt.wantToken.value)
			}
		})
	}
}

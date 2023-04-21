package regexLexer

import (
	"regexp"
)

type Token struct {
	value     string
	tokenType TokenType
}

type TokenType string

const (
	Whitespace         TokenType = "whitespace"
	Number             TokenType = "number"
	Identifier         TokenType = "identifier"
	InferredAssignment TokenType = "inferred_assignment"
	Assignment         TokenType = "assignment"
	Addition           TokenType = "addition"
	Operator           TokenType = "operator"
	Semicolon          TokenType = "semicolon"
	OpenParenthesis    TokenType = "open_parenthesis"
	CloseParenthesis   TokenType = "close_parenthesis"
	OpenCurly          TokenType = "open_curly"
	CloseCurly         TokenType = "close_curly"
	Equal              TokenType = "equal"
	Tab                TokenType = "tab"
	NewLine            TokenType = "new_line"
	Return             TokenType = "return"
	Keyword            TokenType = "keyword"
	Residual           TokenType = "residual"
	Comment            TokenType = "comment"
	String             TokenType = "string"
	Colon              TokenType = "colon"
	Comma              TokenType = "comma"
	Dot                TokenType = "dot"
	Pointer            TokenType = "pointer"
	Reference          TokenType = "reference"
)

var patternsArr = []TokenType{
	Comment,
	String,
	Whitespace,
	Number,
	Identifier,
	Equal,
	InferredAssignment,
	Assignment,
	Colon,
	Addition,
	Operator,
	Semicolon,
	OpenParenthesis,
	CloseParenthesis,
	OpenCurly,
	CloseCurly,
	Tab,
	NewLine,
	Residual,
	Comma,
	Pointer,
	Reference,
	Dot,
	Return,
}

var keywordsArr = []string{
	"package",
	"import",
	"func",
	"int",
	"string",
	"rune",
	"bool",
	"for",
	"range",
	"if",
	"make",
	"struct",
	"type",
	"self",
	"return",
	"var",
	"byte",
	"switch",
	"case",
	"break",
	"continue",
}

var patternsMap = map[TokenType]string{
	Whitespace:         `\s`,
	Number:             `\d`,
	Identifier:         `[\w]`,
	String:             `\"[^"]*\"`,
	Operator:           `<|>|<=|>=`,
	Equal:              `==`,
	InferredAssignment: `:=`,
	Colon:              `:`,
	Assignment:         `=`,
	Addition:           `\+\+`,
	Semicolon:          `;`,
	OpenCurly:          `{`,
	CloseCurly:         `}`,
	OpenParenthesis:    `\(`,
	CloseParenthesis:   `\)`,
	NewLine:            `\n`,
	Tab:                `\t`,
	Residual:           `%`,
	Comment:            `\/\/.*`,
	Comma:              `,`,
	Dot:                `\.`,
	Reference:          `&`,
	Pointer:            `\*`,
	Return:             `\r`,
}

func GetTokens(code string) []Token {
	var tokenArr []Token

	pattern := regexp.MustCompile(`\r|\/\/.*| |[A-z0-9]+|%|==|:=|=|\+\+|\+|\(|\)|<|>|>=|<=|;|{|}|\"[^"]*\"|:|,|\.|\*|&`)
	lexemes := pattern.FindAllString(code, -1)

	for _, lexeme := range lexemes {
		for _, pattern := range patternsArr {
			matchKeyword := false
			for _, keyword := range keywordsArr {
				if keyword == lexeme {
					tokenArr = append(tokenArr, Token{value: lexeme, tokenType: Keyword})
					matchKeyword = true
					break
				}
			}
			if matchKeyword {
				break
			}
			matched, _ := regexp.MatchString(patternsMap[pattern], lexeme)
			if matched {
				tokenArr = append(tokenArr, Token{value: lexeme, tokenType: pattern})
				break
			}
		}
	}
	return tokenArr
}

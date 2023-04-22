package regexLexer

import (
	"fmt"
	"regexp"
	"strings"
)

type Token struct {
	value     string
	tokenType TokenType
}

type TokenType string

const (
	Whitespace          TokenType = "whitespace"
	Number              TokenType = "number"
	Identifier          TokenType = "identifier"
	Assignment          TokenType = "assignment"
	Addition            TokenType = "addition"
	Operators           TokenType = "operator"
	ComparisonOperators TokenType = "comparison_operator"
	Semicolon           TokenType = "semicolon"
	OpenParenthesis     TokenType = "open_parenthesis"
	CloseParenthesis    TokenType = "close_parenthesis"
	OpenCurly           TokenType = "open_curly"
	CloseCurly          TokenType = "close_curly"
	Equal               TokenType = "equal"
	Tab                 TokenType = "tab"
	NewLine             TokenType = "new_line"
	Return              TokenType = "return"
	Keyword             TokenType = "keyword"
	Residual            TokenType = "residual"
	Comment             TokenType = "comment"
	String              TokenType = "string"
	Colon               TokenType = "colon"
	Comma               TokenType = "comma"
	Dot                 TokenType = "dot"
	Reference           TokenType = "reference"
)

var patternsArr = []TokenType{
	Return,
	Comment,
	String,
	Whitespace,
	Number,
	Identifier,
	Equal,
	Assignment,
	Colon,
	Addition,
	Operators,
	Semicolon,
	OpenParenthesis,
	CloseParenthesis,
	OpenCurly,
	CloseCurly,
	Tab,
	NewLine,
	Residual,
	Comma,
	Reference,
	Dot,
}

// "package",
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
	Whitespace:          `\s`,
	Number:              `\d`,
	Identifier:          `[\w]`,
	String:              `\"[^"]*\"`,
	ComparisonOperators: `<|>|<=|>=`,
	Operators:           `+|-|*|/`,
	Equal:               `==`,
	Colon:               `^:$`,
	Assignment:          `^=$`,
	Addition:            `\+\+`,
	Semicolon:           `;`,
	OpenCurly:           `{`,
	CloseCurly:          `}`,
	OpenParenthesis:     `\(`,
	CloseParenthesis:    `\)`,
	NewLine:             `\n`,
	Tab:                 `\t`,
	Residual:            `%`,
	Comment:             `\/\/.*`,
	Comma:               `,`,
	Dot:                 `\.`,
	Reference:           `&`,
	Return:              `\r`,
}

func appendKeywords(extraKeywords []string) {
	for _, v := range extraKeywords {
		keywordsArr = append(keywordsArr, v)
	}
	fmt.Println(keywordsArr)
}

// InferredAssignment:  `:=`,
func appendPatterns(extraPatterns []string) {
	for _, v := range extraPatterns {
		extraPatternArr := strings.Split(v, "0")
		var newToken TokenType = TokenType(extraPatternArr[0])
		newPattern := extraPatternArr[1]
		patternsArr = append(patternsArr, newToken)
		patternsMap[newToken] = newPattern
	}
	fmt.Println(patternsArr)
	fmt.Println(patternsMap)
}

func processFlag(flag string) []string {
	trimmedFlag := strings.TrimSpace(flag)
	flagArray := strings.Split(trimmedFlag, ",")
	return flagArray
}

func GetToken(lexeme string, tokenArr *[]Token) {
	for _, pattern := range patternsArr {
		matchKeyword := false
		for _, keyword := range keywordsArr {
			if keyword == lexeme {
				*tokenArr = append(*tokenArr, Token{value: lexeme, tokenType: Keyword})
				matchKeyword = true
				break
			}
		}
		if matchKeyword {
			break
		}
		matched, _ := regexp.MatchString(patternsMap[pattern], lexeme)
		if matched {
			*tokenArr = append(*tokenArr, Token{value: lexeme, tokenType: pattern})
			break
		}
	}
}
func GetTokens(code string, keywordsFlag string, patternFlag string) []Token {
	var tokenArr []Token

	pattern := regexp.MustCompile(`\r|\/\/.*| |[A-z0-9]+|%|==|:=|^=$|\+\+|\+|\(|\)|<|>|>=|<=|;|{|}|\"[^"]*\"|:|,|\.|\*|&`)
	lexemes := pattern.FindAllString(code, -1)
	if keywordsFlag != "" {
		extraKeywordsArr := processFlag(keywordsFlag)
		appendKeywords(extraKeywordsArr)
	}

	if patternFlag != "" {
		extraPatternArray := processFlag(patternFlag)
		appendPatterns(extraPatternArray)
	}

	for _, lexeme := range lexemes {
		GetToken(lexeme, &tokenArr)
	}
	return tokenArr
}

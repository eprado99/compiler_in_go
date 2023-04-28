package regexLexer

import (
	"encoding/json"
	"regexp"
)

type Lexer struct {
	patternsArray []TokenType
	keywordsArray []string
	patternsMap   map[TokenType]string
}

func (l *Lexer) Load() {
	extraKeywords, extraPatterns, err := l.unmarshal()
	if err == nil {
		l.loadPatterns(extraPatterns)
		l.loadKeywords(extraKeywords)
	}
}

func (l *Lexer) loadPatterns(extraPatterns map[string]string) {
	l.patternsMap = make(map[TokenType]string)
	for k, v := range extraPatterns {
		l.patternsMap[TokenType(k)] = v
		l.patternsArray = append(l.patternsArray, TokenType(k))
	}
}

func (l *Lexer) loadKeywords(extraKeywords []string) {
	l.keywordsArray = append(l.keywordsArray, extraKeywords...)
}

func (l *Lexer) unmarshal() ([]string, map[string]string, error) {
	type lexerJSON struct {
		Keywords []string          `json:"keywords"`
		Patterns map[string]string `json:"patterns"`
	}
	jsonStr := `
	{
		"keywords": [
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
			"continue"
		],
		"patterns": {
			"whitespace": "\\s",
			"number": "\\d",
			"identifier": "[\\w]",
			"string": "\\\"[^\"]*\\\"",
			"comparisonOperators": "<|>|<=|>=",
			"arithmeticOperators": "+|-|*|/",
			"equal": "==",
			"colon": "^:$",
			"assignment": "^=$",
			"addition": "\\+\\+",
			"semicolon": ";",
			"openCurly": "{",
			"closeCurly": "}",
			"openParenthesis": "\\(",
			"closeParenthesis": "\\)",
			"newLine": "\n",
			"tab": "\t",
			"residual": "%",
			"comment": "\/\/.*",
			"comma": ",",
			"dot": "\\.",
			"return": "\\r"
		}
			
	}`

	var lJ lexerJSON
	err := json.Unmarshal([]byte(jsonStr), &lJ)
	if err != nil {
		return nil, nil, err
	}

	return lJ.Keywords, lJ.Patterns, err
}

type Token struct {
	value     string
	tokenType TokenType
}

type TokenType string

const Keyword TokenType = "keyword"

func GetToken(l Lexer, lexeme string, tokenArr *[]Token) {
	for _, pattern := range l.patternsArray {
		matchKeyword := false
		for _, keyword := range l.keywordsArray {
			if keyword == lexeme {
				*tokenArr = append(*tokenArr, Token{value: lexeme, tokenType: Keyword})
				matchKeyword = true
				break
			}
		}
		if matchKeyword {
			break
		}
		matched, _ := regexp.MatchString(l.patternsMap[pattern], lexeme)
		if matched {
			*tokenArr = append(*tokenArr, Token{value: lexeme, tokenType: pattern})
			break
		}
	}
}

func GetTokens(code string, keywordsFlag string, patternFlag string) []Token {
	l := Lexer{}
	var tokenArr []Token
	l.Load()
	pattern := regexp.MustCompile(`\r|\/\/.*| |[A-z0-9]+|%|==|:=|^=$|\+\+|\+|\(|\)|<|>|>=|<=|;|{|}|\"[^"]*\"|:|,|\.|\*|&`)
	lexemes := pattern.FindAllString(code, -1)

	for _, lexeme := range lexemes {
		GetToken(l, lexeme, &tokenArr)
	}
	return tokenArr
}

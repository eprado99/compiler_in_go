package main

// TODO: Use os along Open & Read
import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type token struct {
	value     string
	tokenType string
}

type TokenType string

const (
	Whitespace       TokenType = "whitespace"
	Number           TokenType = "number"
	Identifier       TokenType = "identifier"
	Assignment       TokenType = "assignment"
	Addition         TokenType = "addition"
	Operator         TokenType = "operator"
	Semicolon        TokenType = "semicolon"
	OpenParenthesis  TokenType = "open_parenthesis"
	CloseParenthesis TokenType = "close_parenthesis"
	OpenCurly        TokenType = "open_curly"
	CloseCurly       TokenType = "close_curly"
	Equal            TokenType = "equal"
)

const (
	IDEN         = "IDEN"
	SPACE        = ' '
	EQUALS       = '='
	DOUBLE_QUOTE = '"'
)

var keywords = map[string]string{
	"=":   "EQUALS",
	"var": "VAR",
}
var lexemes = []string{}
var tokens = []token{}

var patterns = map[string]TokenType{
	`\s`:   Whitespace,
	`\d`:   Number,
	`[\w]`: Identifier,
	`<`:    Operator,
	`>`:    Operator,
	`==`:   Equal,
}

func main() {

	file, err := os.Open("hola.txt")
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(file)
	var chars bytes.Buffer

	readingString := false
	for {
		if c, _, err := r.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			checkTokenType(c, &readingString, &chars)
		}
	}

	for _, v := range tokens {
		fmt.Println(v.value, v.tokenType)
	}
}

func isLetter(c rune) bool {
	return 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z' || c == '_'
}

func isDigit(d rune) bool {
	return '0' <= d && d <= '9'
}

func getToken(iden string) string {
	tokenType, ok := keywords[iden]
	if ok {
		return tokenType
	}
	return IDEN
}

func isToken(iden string) bool {
	_, ok := keywords[iden]
	return ok
}

func checkTokenType(c rune, readingString *bool, chars *bytes.Buffer) {
	if isLetter(c) {
		chars.WriteRune(c)
	} else if isDigit(c) {
		tokens = append(tokens, token{value: string(c), tokenType: "NUMBER"})
	} else if c == SPACE {
		lexemes = append(lexemes, chars.String())
		if chars.Len() > 0 {
			tokens = append(tokens, token{value: chars.String(), tokenType: getToken(chars.String())})
		}
		chars.Reset()
	} else if c == EQUALS {
		lexemes = append(lexemes, chars.String())
		if chars.Len() > 0 {
			tokens = append(tokens, token{value: chars.String(), tokenType: getToken(chars.String())})
		}
		chars.Reset()
		lexemes = append(lexemes, string(c))
		tokens = append(tokens, token{value: string(c), tokenType: getToken(string(c))})
	} else if c == DOUBLE_QUOTE || *readingString {
		chars.WriteRune(c)
		if *readingString && c == DOUBLE_QUOTE {
			*readingString = false
			lexemes = append(lexemes, chars.String())
			tokens = append(tokens, token{value: chars.String(), tokenType: "STRING"})
			chars.Reset()
		}
		if !*readingString {
			*readingString = true
		}

	}
}

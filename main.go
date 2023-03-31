package main

// TODO: Use os along Open & Read
import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type token struct {
	value     string
	tokenType string
}

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

func main() {

	file, err := os.Open("hola.txt")
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(file)
	chars := ""
	readingString := false
	for {
		if c, _, err := r.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			if isLetter(c) {
				chars += string(c)

			} else if c == SPACE {
				lexemes = append(lexemes, chars)
				if len(chars) > 0 {
					tokens = append(tokens, token{value: chars, tokenType: getToken(chars)})
				}
				chars = ""
			} else if c == EQUALS {
				lexemes = append(lexemes, chars)
				if len(chars) > 0 {
					tokens = append(tokens, token{value: chars, tokenType: getToken(chars)})
				}
				chars = ""
				lexemes = append(lexemes, string(c))
				tokens = append(tokens, token{value: string(c), tokenType: getToken(string(c))})
			} else if c == DOUBLE_QUOTE || readingString {
				chars += string(c)
				if readingString && c == DOUBLE_QUOTE {
					readingString = false
					lexemes = append(lexemes, chars)
					tokens = append(tokens, token{value: chars, tokenType: getToken(chars)})
					chars = ""
				}
				if !readingString {
					readingString = true
				}

			}
		}
	}

	fmt.Println(lexemes)
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

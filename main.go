package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/eprado99/compiler_in_go/compiler/regexLexer"
)

const (
	keywordsNotProvided = "keyword not provided"
	patternsNotProvided = "pattern not provided"
	keywordsUsage       = "The lexer accepts optional keywords in an attempt to make it language agnostic, the format should be a string and it must be comma-separated, an example of this would be func keyword in Go"
	patternsUsage       = `The lexer accepts additional optional patterns for special language syntax, 
							the format should be a string and it must be comma-separated, 
							an example of this would be the inferred assignment from Go known to be := which would be provided in the following way inferredAssignment:':=', 
							in this case you need to provide the name and pattern separated by a colon and surround the pattern in single quotes`
)

func main() {
	keywordPtr := flag.String("keywords", keywordsNotProvided, keywordsUsage)
	patternPtr := flag.String("patterns", patternsNotProvided, patternsUsage)
	flag.Parse()
	fmt.Println(*keywordPtr)
	fmt.Println(*patternPtr)
	data, err := ioutil.ReadFile("codigo1.txt")
	if err != nil {
		log.Fatal(err)
	}
	keywordsFlag := ""
	patternFlag := ""
	if *keywordPtr != keywordsNotProvided {
		keywordsFlag = *keywordPtr
	}

	if *patternPtr != patternsNotProvided {
		patternFlag = *patternPtr
	}

	tokensArr := regexLexer.GetTokens(string(data), keywordsFlag, patternFlag)
	for k, v := range tokensArr {
		fmt.Println(k, v)
	}
	// fmt.Println(tokensArr)
}

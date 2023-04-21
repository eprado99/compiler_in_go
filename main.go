package main

import (
	"fmt"
	"github.com/eprado99/compiler_in_go/compiler/regexLexer"
	"io/ioutil"
	"log"
)

func main() {

	data, err := ioutil.ReadFile("codigo3.txt")
	if err != nil {
		log.Fatal(err)
	}

	tokensArr := regexLexer.GetTokens(string(data))
	fmt.Println(tokensArr)
}

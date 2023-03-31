package main

// TODO: Use os along Open & Read
import (
	"fmt"
	"io/ioutil"
)

func oldmain() {
	// letters := map[string]bool{
	// 	"a": true, "b": true, "c": true, "d": true, "e": true,
	// }
	file, err := ioutil.ReadFile("hola.txt")
	if err != nil {
		return
	}
	fmt.Println(string(file))
	// cont := 0
	size := len(file)
	fmt.Println(size)
	for i := 0; i < size; i++ {
		fmt.Println(string(file[i]))
		//if(sentence[i] != ' ' && sentence[i+1] == ' ' || sentence[i+1] == '\n' || (i + 2) >= size) {
		//	cont++
		//}
	}
	//fmt.Println(cont)
}

// func isLetter(c rune) bool {
// 	return 'a' <= c && c <= 'z'
// }

// func isDigit(d rune) bool {
// 	return '0' <= d && d <= '9'
// }

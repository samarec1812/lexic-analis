package main

import (
	"fmt"
	"github.com/samarec1812/lexic-analis/analysator"
	"io/ioutil"
	_ "regexp"
	"strconv"
	"strings"
)

//func SplitToken(word string) {
//	for _, sym := range word {
//		if unicode.IsLetter(sym)
//	}
//}

func main() {
	// 2 sposob
	content, err := ioutil.ReadFile("input8.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	prog := strings.ToLower(string(content))
	fmt.Println(prog)
	fmt.Println()
	tokens, tokens2 := analysator.SplitText(prog)
	for idx, token := range tokens {
		if token == "\n" {
			fmt.Print("Перенос строки : ")
		}
		fmt.Println(strconv.Itoa(idx) + " : " + token)
	}
	check := analysator.Checker(tokens)
	fmt.Println(check)
	if !check {
		fmt.Println(analysator.GetErrorLine(tokens2))
	}
	// tokensLength := len(tokens)
	fmt.Println(tokens)
}

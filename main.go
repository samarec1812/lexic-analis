package main

import (
	"fmt"
	"github.com/samarec1812/lexic-analis/analysator"
	"io/ioutil"
	"os"
	_ "regexp"
	"strconv"
	"strings"
)


func main() {

	content, err := ioutil.ReadFile("input10.txt")
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
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	finallyStr :=  prog + "\n"
	finallyStr += "---------------------------\n"
	finallyStr += strconv.FormatBool(check)
	finallyStr += "\n"
	if !check {
		finallyStr += strconv.Itoa(analysator.GetErrorLine(tokens2))
	}

	file.WriteString(finallyStr)


}

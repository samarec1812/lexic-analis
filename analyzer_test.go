package main

import (
	"fmt"
	"github.com/samarec1812/lexic-analis/analysator"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func wrapper(fileName string) []string {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return []string{}
	}
	defer f.Close()

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return []string{}
	}

	prog := strings.ToLower(string(content))
	tokens, _ := analysator.SplitText(prog)
	prog = ""
	content = nil
	return tokens
}

func TestChecker_4(t *testing.T) {
	if analysator.Checker(wrapper("input4.txt")) != true {
		t.Error("TEST 4: false != true")
	}
}

func TestChecker_1(t *testing.T) {

	if analysator.Checker(wrapper("input.txt")) != false {
		t.Error("TEST 1: true != false")
	}
}

func TestChecker_2(t *testing.T) {
	if analysator.Checker(wrapper("input2.txt")) != false {
		t.Error("TEST 2: false != true")
	}
}

func TestChecker_3(t *testing.T) {
	if analysator.Checker(wrapper("input3.txt")) != false {
		t.Error("TEST 3: false != true")
	}
}

func TestChecker_5(t *testing.T) {
	if analysator.Checker(wrapper("input5.txt")) != false {
		t.Error("TEST 5: false != true")
	}
}


func TestChecker_9(t *testing.T) {
	if analysator.Checker(wrapper("input9.txt")) != false {
		t.Error("TEST 9: false != true")
	}
}


func TestChecker_6(t *testing.T) {
	if analysator.Checker(wrapper("input6.txt")) != true {
		t.Error("TEST 6: false != true")
	}
}

func TestChecker_7(t *testing.T) {
	if analysator.Checker(wrapper("input7.txt")) != true {
		t.Error("TEST 7: false != true")
	}
}

func TestChecker_8(t *testing.T) {
	if analysator.Checker(wrapper("input8.txt")) != true {
		t.Error("TEST 8: false != true")
	}
}


func TestChecker_10(t *testing.T) {
	if analysator.Checker(wrapper("input10.txt")) != true {
		t.Error("TEST 10: false != true")
	}
}
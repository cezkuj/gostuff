package main

import (
	"testing"
        "fmt"
)

type myString string

func TestSubtypes(*testing.T) {
	var myS myString
	var commonS string
        commonS = commonToCommon(string(myS))
        fmt.Println(commonS)

}

func myToMy(s myString) myString {
	return s
}

func myToCommon(s myString) string {
	return string(s)
}

func commonToMy(s string) myString {
	return myString(s)
}

func commonToCommon(s string) string {
	return s
}

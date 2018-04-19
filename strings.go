package main

import (
	"fmt"
	"strings"
)

func test_strings() {
	name := "example.com"
	fmt.Println(strings.Trim(name, "."))
	fmt.Println(strings.Contains(strings.Trim(name, "."), "."))
	fmt.Println(strings.Contains("abc", "b"))
	if !strings.Contains(strings.Trim(name, "."), ".") {
		fmt.Println("acme/autocert: server name component count invalid")
	}
}

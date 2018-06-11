package main

import (
	"fmt"
	"testing"
)

type A struct {
	number int
}
type B struct {
}

type C struct {
	A
	B
}

func (a A) print() {
	fmt.Println(a, "a")
}

func (b B) print() {
	fmt.Println(b, "b")
}

func TestInh(*testing.T) {
	a := A{2}
	c := C{A:a}
	fmt.Println(a)
	fmt.Println(c)
        fmt.Println(c.number)
}

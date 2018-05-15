package main

import (
	"fmt"
	"testing"
)

type A struct {
	number int
}
type B struct {
        number int
}

type C struct {
	A
        B
	number int
}
func (a A) print() {
 fmt.Println(a, "a")
}

func (b B) print() {
 fmt.Println(b, "b")
}


func TestInh(*testing.T) {
        a := A{}
        c := C{}
	fmt.Println(a)
	fmt.Println(c)
        a.print()
        c.A.print()
        c.print()

}

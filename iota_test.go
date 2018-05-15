package main

import (
	"fmt"
	"testing"
)

func TestIota(*testing.T) {
	const (
		a = 1
		b = iota
		c
	)
        const (
d = 1 - iota
e 
f


)
    	fmt.Println(a, b, c)
        fmt.Println(d, e ,f)
}

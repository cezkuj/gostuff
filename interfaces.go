package main

import "fmt"

type a interface {
	doSth()
}

type B struct {
	b int
}

func (b *B) doSth() {
	fmt.Println("4 exec")
	b.b = 4
}

func test_interfaces() {
	b := B{b: 1}
	b_p := &B{b: 2}
	fmt.Println(b.b, b_p.b)
	b.doSth()
	b_p.doSth()
	fmt.Println(b.b, b_p.b)
}

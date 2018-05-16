package main

import (
	"fmt"
	"testing"
)

func appender(slice *[]int) {
	(*slice)[0] = 3
	*slice = append(*slice, 10)
}

func TestSlices(T *testing.T) {
	a := []int{}
	//a := make([]int, 0) //same as above
	fmt.Println(a)
	a = append(a, 5)
	appender(&a)
	fmt.Println(a)
}

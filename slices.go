package main

import "fmt"

func test_slices() {
	//a := []int{}
	a := make([]int, 0) //same as above
	fmt.Println(a)
	a = append(a, 5)
	fmt.Println(a)
}

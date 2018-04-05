package main

import "fmt"

var b int = 2
func init(){
  fmt.Println("hello init")
}

func changer() {
  b += 2
}
func scope(){
  a := 1
  fmt.Println(a, b)
  changer()
  fmt.Println(a, b)
}

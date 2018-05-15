package main

import (
	"log"
	"testing"
)

func TestChannelsRange(*testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, number := range numbers {
		go func(number int) {
			log.Println(number)
		}(number)
	}

}

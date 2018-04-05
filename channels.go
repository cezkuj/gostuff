package main

import (
	"fmt"
	"sync"
	"time"
)

type Differ struct {
	id   int
	data string
}

func add_to_channel(differ Differ, c chan Differ, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 20; i++ {
		c <- differ
		time.Sleep(time.Duration(differ.id) * time.Second)
	}
}

func chanells() {
	c := make(chan Differ)
	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go add_to_channel(Differ{i, "abc"}, c, wg)
	}
	go func() {
		wg.Wait()
		close(c)
	}()
	for j := range c {
		fmt.Println(j)
	}
}

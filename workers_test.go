package main

import (
	"fmt"
	"testing"
	"time"
)

type Worker struct {
	id int
}

func (w Worker) work(r int) {
	time.Sleep(time.Second * 2)
	fmt.Println(w.id, r)
}
func TestWorkers(*testing.T) {
	requests := make(chan int)
	workers := make(chan Worker, 5)
	for i := 0; i < 5; i++ {
		workers <- Worker{i}
	}
	go func() {
		for {
			requests <- 1
		}
	}()
	for request := range requests {
		worker := <-workers
		go func() {
			worker.work(request)
			workers <- worker
		}()
	}

}

package main

import (
	"fmt"
)

type MyWaitGroup struct {
	counter int
	done    chan struct{}
}

var slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var results = []int{}

func main() {
	var mwg MyWaitGroup
	mwg.done = make(chan struct{})
	intCh := make(chan int)

	for _, i := range slice {
		mwg.Add(1)
		go func(num int) {
			defer mwg.Done()
			intCh <- num
		}(i)
	}

	go func() {
		mwg.Wait()
		close(intCh)
	}()

	for n := range intCh {
		results = append(results, n)
	}
	fmt.Println(results)
}

func (mwg *MyWaitGroup) Add(num int) {
	mwg.counter += num
	if mwg.counter == 0 {
		close(mwg.done)
	}
}

func (mwg *MyWaitGroup) Done() {
	mwg.Add(-1)
}

func (wg *MyWaitGroup) Wait() {
	<-wg.done
}

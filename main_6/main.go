package main

import (
	"fmt"
	"sync"
)

var slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var results = []int{}
var wg sync.WaitGroup

func main() {
	intCh := make(chan int)

	for _, i := range slice {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			intCh <- num
		}(i)
	}

	go func() {
		wg.Wait()
		close(intCh)
	}()

	for n := range intCh {
		results = append(results, n)
	}
	fmt.Println(results)
}

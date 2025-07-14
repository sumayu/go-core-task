package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func merge(slice ...chan int) chan int {
	result := make(chan int)

	for _, ch := range slice {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range ch {
				result <- j
			}
		}()
	}
	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}

func main() {
	chOne := make(chan int)
	chTwo := make(chan int)
	chFree := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			chOne <- i
		}
		close(chOne)
	}()

	go func() {
		for i := 10; i < 20; i++ {
			chTwo <- i
		}
		close(chTwo)
	}()

	go func() {
		for i := 20; i < 30; i++ {
			chFree <- i
		}
		close(chFree)
	}()

	for num := range merge(chOne, chTwo, chFree) {
		fmt.Println(num)
	}
}

package main

import "fmt"

func main() {
	chUint8 := make(chan uint8)
	chUint16 := make(chan float64)
	numbers := generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	stage := redesign(numbers, chUint8)
	stagetwo := redesignFloat(stage, chUint16)
	for num := range stagetwo {
		fmt.Printf("%f\n", num)
	}

}

func generate(nums ...int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, n := range nums {
			ch <- n
		}
	}()
	return ch
}

func redesign(nums <-chan int, ch chan uint8) <-chan uint8 {
	go func() {
		defer close(ch)
		for n := range nums {
			ch <- uint8(n)
		}
	}()
	return ch
}

func redesignFloat(nums <-chan uint8, ch chan float64) <-chan float64 {
	go func() {
		defer close(ch)
		for n := range nums {
			num := float64(n)
			ch <- num * num * num
		}
	}()
	return ch
}

package main

import (
	"fmt"
	"go_core_task/main_2/handler"
	"math/rand"
)

func main() {
	originalSlice := make([]int, 10)

	for i := range originalSlice {
		originalSlice[i] = rand.Intn(100)
	}

	fmt.Println("Случайный слайс:", originalSlice)

	chetkiy := handler.SliceExample(originalSlice)
	fmt.Println("четкий слайз", chetkiy)

	addslice := handler.AddElements(originalSlice, 1)
	fmt.Println("добавленный элимент", addslice)

	copeclice := handler.CopySlice(originalSlice)
	fmt.Println("копироемуй элимент", copeclice)

	deleteslice := handler.RemoveElement(originalSlice, 4)
	fmt.Println("удаленный слайз", deleteslice)

}

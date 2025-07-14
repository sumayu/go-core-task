package main

import "fmt"

func intersection(slice1, slice2 []int) (bool, []int) {
	present := make(map[int]bool)
	for _, item := range slice1 {
		present[item] = true
	}

	var diff []int
	found := false

	for _, item := range slice2 {
		if present[item] {
			diff = append(diff, item)
			present[item] = false
			found = true
		}
	}
	return found, diff
}

func main() {
	slice1 := []int{65, 3, 58, 678, 64}
	slice2 := []int{64, 2, 3, 43}
	fmt.Println(intersection(slice1, slice2))
}

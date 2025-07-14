package main

import "fmt"

func Difference(slice1, slice2 []string) []string {
	present := make(map[string]bool)
	for _, item := range slice2 {
		present[item] = true
	}

	var diff []string
	for _, item := range slice1 {
		if !present[item] {
			diff = append(diff, item)
		}
	}

	return diff
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	fmt.Println(Difference(slice1, slice2))
}

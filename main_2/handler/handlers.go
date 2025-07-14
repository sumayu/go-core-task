package handler

func SliceExample(slice []int) []int {
	originalSlice := make([]int, 0)
	for i := range slice {
		if slice[i]%2 == 0 {
			originalSlice = append(originalSlice, slice[i])
		}
	}
	return originalSlice
}

func AddElements(slice []int, n int) []int {
	newSlice := append(slice, n)
	return newSlice
}

func CopySlice(slice []int) []int {
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	return newSlice
}

func RemoveElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

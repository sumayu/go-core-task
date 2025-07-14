package main

import "fmt"

func add(StringIntMap map[string]int, key string, value int) map[string]int {
	StringIntMap[key] = value
	return StringIntMap
}
func remove(StringIntMap map[string]int, key string) map[string]int {
	delete(StringIntMap, key)
	return StringIntMap
}
func Copy(StringIntMap map[string]int) map[string]int {
	newMap := make(map[string]int, len(StringIntMap))
	for k, v := range StringIntMap {
		newMap[k] = v
	}
	return newMap
}
func Exists(StringIntMap map[string]int, key string) bool {
	value, exists := StringIntMap[key]
	if exists {
		fmt.Println("banana exists with value", value)
	}
	return exists
}

func main() {
	StringIntMap := make(map[string]int)
	add(StringIntMap, "apple", 5)
	Copy(StringIntMap)
	Exists(StringIntMap, "apple")
	remove(StringIntMap, "apple")
}

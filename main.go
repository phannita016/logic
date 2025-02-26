package main

import "fmt"

func mergeArrays(arr1, arr2 []string) ([]string, []string) {
	uniqueMap := make(map[string]bool)
	union := []string{}
	intersect := []string{}

	for i := 0; i < len(arr1); i++ {
		if !uniqueMap[arr1[i]] {
			uniqueMap[arr1[i]] = true
			union = append(union, arr1[i])
		}
	}

	for i := 0; i < len(arr2); i++ {
		if !uniqueMap[arr2[i]] {
			uniqueMap[arr2[i]] = true
			union = append(union, arr2[i])
		} else {
			uniqueMap[arr2[i]] = false
		}
	}

	for key, exists := range uniqueMap {
		if exists {
			intersect = append(intersect, key)
		}
	}

	return union, intersect
}

func main() {
	arr1 := []string{"a", "b", "c"}
	arr2 := []string{"b", "c", "d"}
	union, intersect := mergeArrays(arr1, arr2)

	fmt.Println("Union:", union)         // [a b c d]
	fmt.Println("Intersect:", intersect) // [a d]
}

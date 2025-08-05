package main

import "fmt"

// you can use pointers to modify the underlying slice

func appendSafely(s *[]int, value int) {
	*s = append(*s, value)
}

// you can return and reassign to modify the length and capacity
func appendToSlice(s []int) []int {
	return append(s, 4)
}

// pros return the modified slice
func addElements(s []int, element int) []int {
	return append(s, element)
}

// this returns the modified slice in result by taking in the length of the original slice (s)
func filterEvens(s []int) []int {
	result := make([]int, 0, len(s)) // pre-allocate with capacity
	for _, v := range s {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	nums := []int{1, 2, 3}
	appendSafely(&nums, 4)
	fmt.Println(nums, len(nums), cap(nums))

	nums2 := []int{1, 2, 3}
	nums2 = appendToSlice(nums2)
	fmt.Println(nums2)
	fmt.Println(filterEvens(nums2))
	fmt.Println(nums2)
}

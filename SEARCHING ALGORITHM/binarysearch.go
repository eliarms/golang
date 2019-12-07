// Binary Search in Golang
package main

import "fmt"

func binarySearch(needle int, numbers []int) bool {

	low := 0
	high := len(numbers) - 1

	for low <= high {
		median := (low + high) / 2

		if numbers[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(numbers) || numbers[low] != needle {
		return false
	}

	return true
}

func main() {
	numbers := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(binarySearch(1, numbers))
}

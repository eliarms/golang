package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{1, 3, 4, 5, 9, 10, 2, 6}
	tobeSorted := sort.IntSlice(numbers)
	sort.Sort(sort.Reverse(tobeSorted))
	fmt.Println(tobeSorted)
}

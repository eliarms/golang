package main

import "fmt"

func main() {
	items1 := []int{1, 4}
	items2 := []int{2, 3}
	result := append(items1, items2...)
	fmt.Println(result)
}

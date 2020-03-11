package main

import "fmt"

func main() {
	numbers := []int{1, 3, 4, 5, 9, 10, 2, 6}
	for _, value := range numbers {
		//fmt.Printf("Index: %v and Value: %v\n", index, value)
		fmt.Println(value)
	}

}

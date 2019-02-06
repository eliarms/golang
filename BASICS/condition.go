package main

import "fmt"

func main() {
	x := 10
	if x > 5 {
		fmt.Println("X is big")
	}
	if x > 100 {
		fmt.Println("X is very big")
	} else {
		fmt.Println("X is not that big")
	}
	if x > 5 && x < 20 {
		fmt.Println("X is just right")
	}
	if x < 20 || x > 30 {
		fmt.Println("X is out of range ")
	}
}

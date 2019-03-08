package main

import "fmt"

func main() {
	cards := []string{"Demo1", "Demo2"}
	cards = append(cards, "Demo3")
	fmt.Println(cards)
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

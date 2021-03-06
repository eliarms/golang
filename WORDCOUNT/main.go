package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `DevOps DevOps is the combination of cultural philosophies, practices, and tools that increases an organization’s ability to deliver applications and services at high velocity:
			 evolving and improving products at a faster pace than organizations using traditional software development and infrastructure management processes. This speed enables organizations to better serve
			  their customers and compete more effectively in the market.`

	words := strings.Fields(text)
	counts := map[string]int{} //Empty map
	for _, word := range words {
		counts[strings.ToLower(word)]++
	}
	//fmt.Println(counts)
	for w, t := range counts {
		fmt.Println("Total count for", w, "is", t)
	}
}

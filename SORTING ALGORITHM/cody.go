package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "We test coders. Give us a try?"
	max := findWordMax(s) // 9
	fmt.Println(max)

}
func findWordMax(s string) int {
	numberOfWords, maximumWords := 0, 0
	inWord := false
	for _, word := range s {
		switch word {
		case '.', '?', '!':
			inWord = false
			if maximumWords < numberOfWords {
				numberOfWords = numberOfWords
			}
			numberOfWords = 0
		default:
			if unicode.IsSpace(word) {
				inWord = false
			} else if inWord == false {
				inWord = true
				numberOfWords++
			}
		}
		if maximumWords < numberOfWords {
			maximumWords = numberOfWords
		}
	}
	return maximumWords
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	str := []string{"Golang", "is", "an", "awesome", "language"}
	for _, v := range str {
		fmt.Println(v)

	}
	sortedList := sort.StringSlice(str)
	sortedList.Sort()
	fmt.Println(sortedList)
	index := sortedList.Search("awesome")
	fmt.Println(index)

}

package main

import "fmt"

func main() {
	map1 := map[string]int{
		"Michael": 10,
		"Jessica": 20,
		"Tarik":   33,
		"Jon":     22,
	}
	fmt.Println(map1)

	map2 := map[string]int{
		"Lord":  12,
		"Jes":   2,
		"Tik":   3,
		"Jonas": 21,
	}
	for key, value := range map2 {
		map1[key] = value
	}
	fmt.Println(map1)
}

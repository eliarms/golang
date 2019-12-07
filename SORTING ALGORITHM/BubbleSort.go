package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	bubleSort(generateRandomSlice(5))

}

func generateRandomSlice(size int) []int {
	records := make([]int, size, size)
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	for i := 0; i < size; i++ {
		records[i] = r.Intn(50)
	}
	return records

}

func bubleSort(records []int) {

	sort.Slice(records, func(i, j int) bool {
		return records[i] < records[j]
	})
	for _, v := range records {
		fmt.Println(v)
	}

}

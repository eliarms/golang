package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// Interval Definition for an interval.
type Interval struct {
	Start int
	End   int
}

func main() {
	var x string = `4
	0 1
	2 4
	6 7
	3 5`

	//its := []Interval{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	//fmt.Println(merge(its))

	//scanner := bufio.NewScanner(os.Stdin)
	//for scanner.Scan() {
	//	fmt.Println(scanner.Text())
	//}

	//if err := scanner.Err(); err != nil {
	//log.Println(err)
	//}
	for _, line := range strings.Split(strings.TrimSuffix(x, "\n"), "\n") {
		fmt.Println(line)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}

func merge(its []Interval) []Interval {
	if len(its) <= 1 {

		return its
	}

	sort.SliceStable(its, func(i, j int) bool {
		a, b := its[i], its[j]
		return a.Start < b.Start || (a.Start == b.Start && a.End > b.End)
	})
	res := []Interval{}
	cur := its[0]
	for _, p := range its {
		if p.Start > cur.End {
			res = append(res, cur)
			cur = p
		} else {
			if p.End > cur.End {
				cur.End = p.End
			}
		}
	}
	res = append(res, cur)
	return res
}

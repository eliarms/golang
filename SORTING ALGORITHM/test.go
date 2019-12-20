package main

import (
	"fmt"
	"sort"
)

// Interval Definition for an interval.
type Interval struct {
	Start int
	End   int
}

func main() {

	its := []Interval{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(its))
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

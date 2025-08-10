package main

import "sort"

func hIndex(citations []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))
	for i, c := range citations {
		if c <= i {
			return i
		}
	}

	return len(citations)
}

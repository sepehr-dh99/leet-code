package main

import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	boats := 0
	start := 0
	end := len(people) - 1

	for start < end {
		if people[start]+people[end] <= limit {
			boats++
			start++
		} else {
			boats++
		}
		end--
	}

	if start == end {
		boats++
	}

	return boats
}

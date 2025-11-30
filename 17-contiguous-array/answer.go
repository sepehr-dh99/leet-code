package main

func findMaxLength(nums []int) int {
	m := make(map[int]int)
	m[0] = -1 // prefix sum zero before starting

	maxLen := 0
	prefix := 0

	for i, v := range nums {
		if v == 0 {
			prefix -= 1
		} else {
			prefix += 1
		}

		if idx, exist := m[prefix]; exist {
			if i-idx > maxLen {
				maxLen = i - idx
			}
		} else {
			m[prefix] = i
		}
	}

	return maxLen
}

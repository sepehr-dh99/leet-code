package main

func lengthOfLongestSubstring(s string) int {
	lastIndex := make(map[rune]int)
	maxLen := 0
	start := 0

	for i, ch := range s {
		if lastPos, ok := lastIndex[ch]; ok && lastPos >= start {
			start = lastPos + 1
		}
		lastIndex[ch] = i
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
	}

	return maxLen
}

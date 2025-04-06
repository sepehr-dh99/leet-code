package main

func canJump(nums []int) bool {
	jumps := 0

	for _, num := range nums {
		if jumps < 0 {
			return false
		} else if num > jumps {
			jumps = num
		}

		jumps -= 1
	}

	return true
}

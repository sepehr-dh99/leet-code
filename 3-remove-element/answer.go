package main

func removeElement(nums []int, val int) int {
	k := 0

	for index, numb := range nums {
		if numb != val {
			nums[k] = nums[index]
			k += 1
		}
	}

	return k
}

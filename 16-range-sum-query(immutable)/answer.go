package main

type NumArray []int

func Constructor(nums []int) NumArray {
	var numArray NumArray = nums

	for i := 1; i < len(nums); i++ {
		numArray[i] += numArray[i-1]
	}

	return numArray
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return (*this)[right]
	}

	return (*this)[right] - (*this)[left-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

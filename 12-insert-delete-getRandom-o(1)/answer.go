package main

import (
	"math/rand"
)

type RandomizedSet struct {
	nums  []int
	index map[int]int // value -> index in nums
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		nums:  []int{},
		index: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, exists := this.index[val]; exists {
		return false
	}
	this.nums = append(this.nums, val)
	this.index[val] = len(this.nums) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, exists := this.index[val]
	if !exists {
		return false
	}
	lastVal := this.nums[len(this.nums)-1]
	this.nums[idx] = lastVal
	this.index[lastVal] = idx
	this.nums = this.nums[:len(this.nums)-1]
	delete(this.index, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

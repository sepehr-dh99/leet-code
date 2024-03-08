package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	index := m + n - 1
	m--
	n--

	for m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[index] = nums1[m]
			m--
		} else {
			nums1[index] = nums2[n]
			n--
		}
		index--
	}

	for n >= 0 {
		nums1[index] = nums2[n]
		n--
		index--
	}
}

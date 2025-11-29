class NumArray:

    def __init__(self, nums: list[int]):
        self.num_array = nums

        for i in range(1, len(nums)):
            self.num_array[i] += self.num_array[i-1]

    def sumRange(self, left: int, right: int) -> int:
        if left == 0:
            return self.num_array[right]
        else:
            return self.num_array[right] - self.num_array[left - 1]


# Your NumArray object will be instantiated and called as such:
# obj = NumArray(nums)
# param_1 = obj.sumRange(left,right)
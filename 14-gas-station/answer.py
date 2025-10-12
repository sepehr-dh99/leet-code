class Solution:
    def canCompleteCircuit(self, gas: List[int], cost: List[int]) -> int:
        total_tank = 0
        curr_tank = 0
        start = 0

        for i in range(len(gas)):
            diff = gas[i] - cost[i]
            total_tank += diff
            curr_tank += diff

            if curr_tank < 0:
                start = i + 1
                curr_tank = 0

        if total_tank < 0:
            return -1
        return start
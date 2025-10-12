package main

func canCompleteCircuit(gas []int, cost []int) int {
	totalTank := 0
	currTank := 0
	start := 0

	for i := 0; i < len(gas); i++ {
		diff := gas[i] - cost[i]
		totalTank += diff
		currTank += diff

		if currTank < 0 {
			start = i + 1
			currTank = 0
		}
	}

	if totalTank < 0 {
		return -1
	}
	return start
}

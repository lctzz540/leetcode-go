package main

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	numRows := len(matrix)
	numCols := len(matrix[0])
	count := 0

	for left := 0; left < numCols; left++ {
		rowSum := make([]int, numRows)
		for right := left; right < numCols; right++ {
			for row := 0; row < numRows; row++ {
				rowSum[row] += matrix[row][right]
			}
			count += subarraySum(rowSum, target)
		}
	}

	return count
}

func subarraySum(nums []int, k int) int {
	count := 0
	prefixSum := make(map[int]int)
	prefixSum[0] = 1
	currSum := 0

	for _, num := range nums {
		currSum += num
		if prevCount, exists := prefixSum[currSum-k]; exists {
			count += prevCount
		}
		prefixSum[currSum]++
	}

	return count
}

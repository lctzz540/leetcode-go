// Problem 848
package main

func countElementsLessThanOrEqual(arr []int, x int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2
		if arr[mid] <= x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	n := len(obstacles)
	dp := make([]int, n)
	tails := make([]int, n)
	tails[0] = obstacles[0]
	dp[0] = 1
	len := 1

	for i := 1; i < n; i++ {
		if obstacles[i] >= tails[len-1] {
			tails[len] = obstacles[i]
			dp[i] = len + 1
			len++
		} else {
			j := countElementsLessThanOrEqual(tails[:len], obstacles[i])
			tails[j] = obstacles[i]
			dp[i] = j + 1
		}
	}

	return dp
}

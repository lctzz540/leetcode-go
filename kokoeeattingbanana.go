// Problem 875
package main

import "fmt"

func canEatAll(piles []int, k, h int) bool {
	hours := 0
	for _, p := range piles {
		hours += (p + k - 1) / k
	}
	return hours <= h
}

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, 1_000_000_000
	for left < right {
		mid := (left + right) / 2
		if canEatAll(piles, mid, h) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	piles := []int{30, 11, 23, 4, 20}
	h := 5
	fmt.Print(minEatingSpeed(piles, h))
}

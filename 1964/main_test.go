package main

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

func TestLongestObstacleCourseAtEachPosition(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected []int
	}{
		{[]int{1, 2, 3, 2}, []int{1, 2, 3, 3}},
		{[]int{2, 2, 1}, []int{1, 2, 1}},
		{[]int{3, 1, 5, 6, 4, 2}, []int{1, 1, 2, 3, 2, 2}},
		{[]int{5, 1, 5, 5, 1, 3, 4, 5, 1, 4}, []int{1, 1, 2, 3, 2, 3, 4, 5, 3, 5}},
		{[]int{5, 3, 4, 4, 4, 2, 1, 1, 4, 1}, []int{1, 1, 2, 3, 4, 1, 1, 2, 5, 3}},
	}

	for number, tc := range testCases {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		memAllocated := m.Alloc
		result := longestObstacleCourseAtEachPosition(tc.nums)
		runtime.ReadMemStats(&m)
		memAllocated = m.Alloc - memAllocated

		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v \n\t  but got %v\n\t for nums %v", tc.expected, result, tc.nums)
		}

		if !t.Failed() {
			fmt.Printf("Testcase %d PASS, Memory usage: %d bytes\n", number+1, memAllocated)
		}
	}
}

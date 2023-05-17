package main

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

func TestNumSubmatrixSumTarget(t *testing.T) {
	testCases := []struct {
		matrix   [][]int
		target   int
		expected int
	}{
		{[][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}}, 0, 4},
		{[][]int{{1, -1}, {-1, 1}}, 0, 5},
		{[][]int{{904}}, 0, 0},
	}

	for number, tc := range testCases {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		memAllocated := m.Alloc
		result := numSubmatrixSumTarget(tc.matrix, tc.target)
		runtime.ReadMemStats(&m)
		memAllocated = m.Alloc - memAllocated

		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %d \n\t  but got %v\n\t for matrix is %v and target is %d", tc.expected, result, tc.matrix, tc.target)
		}

		if !t.Failed() {
			fmt.Printf("Testcase %d PASS, Memory usage: %d bytes\n", number+1, memAllocated)
		}
	}
}

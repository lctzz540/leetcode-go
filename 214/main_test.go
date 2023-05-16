package main

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

func TestShortestPalindrome(t *testing.T) {
	testCases := []struct {
		s        string
		expected string
	}{
		{"aacecaaa", "aaacecaaa"},
		{"abcd", "dcbabcd"},
	}

	for number, tc := range testCases {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		memAllocated := m.Alloc
		result := shortestPalindrome(tc.s)
		runtime.ReadMemStats(&m)
		memAllocated = m.Alloc - memAllocated

		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v \n\t  but got %v\n\t for nums %v", tc.expected, result, tc.s)
		}

		if !t.Failed() {
			fmt.Printf("Testcase %d PASS, Memory usage: %d bytes\n", number+1, memAllocated)
		}
	}
}

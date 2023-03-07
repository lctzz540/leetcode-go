// Problem 23
package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var arr []int
	for _, listElement := range lists {
		node := listElement
		for node != nil {
			arr = append(arr, node.Val)
			node = node.Next
		}
	}
	sort.Ints(arr)

	var LNout *ListNode
	if len(arr) > 0 {
		LNout = &ListNode{Val: arr[0]}
		node := LNout
		for i := 1; i < len(arr); i++ {
			node.Next = &ListNode{Val: arr[i]}
			node = node.Next
		}
	}
	return LNout
}

func main() {
	lists := []*ListNode{
		{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}},
		{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}},
		{Val: 2, Next: &ListNode{Val: 6, Next: nil}},
	}
	fmt.Print(mergeKLists(lists))
}

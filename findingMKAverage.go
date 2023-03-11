// Problem 1825
package main

import (
	"math/rand"
	"sort"
)

type MKAverage struct {
	L, M, R Range
	m, k    int
	q       []int
}

type Range struct {
	s   []int
	sum int
}

func (r *Range) insert(x int) {
	i := sort.SearchInts(r.s, x)
	r.s = append(r.s, 0)
	copy(r.s[i+1:], r.s[i:])
	r.s[i] = x
	r.sum += x
}

func (r *Range) remove(x int) {
	i := sort.SearchInts(r.s, x)
	r.s = append(r.s[:i], r.s[i+1:]...)
	r.sum -= x
}

func Constructor(m int, k int) MKAverage {
	return MKAverage{
		m: m,
		k: k,
	}
}

func (mk *MKAverage) AddElement(num int) {
	mk.q = append(mk.q, num)
	n := len(mk.q)
	if n < mk.m {
		return
	}
	if n == mk.m {
		qq := make([]int, mk.m)
		copy(qq, mk.q)
		sort.Ints(qq)
		for i := 0; i < mk.k; i++ {
			mk.L.insert(qq[i])
		}
		for i := mk.k; i < mk.m-mk.k; i++ {
			mk.M.insert(qq[i])
		}
		for i := mk.m - mk.k; i < mk.m; i++ {
			mk.R.insert(qq[i])
		}
	}
	if n > mk.m {
		mk.M.insert(num)
		x := mk.L.s[len(mk.L.s)-1]
		y := mk.M.s[0]
		if x > y {
			mk.L.remove(x)
			mk.M.remove(y)
			mk.L.insert(y)
			mk.M.insert(x)
		}
		x = mk.M.s[len(mk.M.s)-1]
		y = mk.R.s[0]
		if x > y {
			mk.M.remove(x)
			mk.R.remove(y)
			mk.M.insert(y)
			mk.R.insert(x)
		}
		invalid := mk.q[n-mk.m-1]
		if i := sort.SearchInts(mk.M.s, invalid); i < len(mk.M.s) && mk.M.s[i] == invalid {
			mk.M.remove(invalid)
		} else if i := sort.SearchInts(mk.L.s, invalid); i < len(mk.L.s) && mk.L.s[i] == invalid {
			mk.L.remove(invalid)
			x := mk.M.s[0]
			mk.L.insert(x)
			mk.M.remove(x)
		} else {
			mk.R.remove(invalid)
			x := mk.M.s[len(mk.M.s)-1]
			mk.R.insert(x)
			mk.M.remove(x)
		}
	}
}

func (mk *MKAverage) CalculateMKAverage() int {
	if len(mk.q) < mk.m {
		return -1
	}
	return mk.M.sum / len(mk.M.s)
}

/**
 * Your MKAverage object will be instantiated and called as such:
 * obj := Constructor(m, k);
 * obj.AddElement(num);
 * param_2 := obj.CalculateMKAverage();
 */
func main() {
	mk := Constructor(100, 50)

	for i := 0; i < 100000; i++ {
		if rand.Intn(2) == 0 {
			mk.AddElement(rand.Intn(100))
		} else {
			mk.CalculateMKAverage()
		}
	}

}

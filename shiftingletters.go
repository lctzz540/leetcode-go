// Problem 848
package main

import (
	"fmt"
	"strings"
)

func shiftingLetters(s string, shifts []int) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	runestr := []rune(s)
	shiftSum := 0

	for i := len(runestr) - 1; i >= 0; i-- {
		shiftSum += shifts[i]
		shiftValue := shiftSum % 26
		idx := strings.IndexRune(alphabet, runestr[i])
		runestr[i] = []rune(alphabet)[(idx+shiftValue)%26]
	}

	return string(runestr)
}

func main() {
	fmt.Println(shiftingLetters("bad", []int{10, 20, 30}))
}

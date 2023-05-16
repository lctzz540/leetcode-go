package main

func shortestPalindrome(s string) string {
	n := len(s)
	reverse := reverseString(s)

	for i := 0; i < n; i++ {
		if s[:n-i] == reverse[i:] {
			return reverse[:i] + s
		}
	}

	return ""
}

func reverseString(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

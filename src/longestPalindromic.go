package main

import "fmt"

// Given a string s, return the longest palindromic substring in s.
// palindrome is a word that reads the same forward as backward e.g. rotator, anna, racecar

// Difficulty: Medium

// Constraints:

// 1 <= s.length <= 1000
// s consist of only digits and English letters.

func longestPalindrome(s string) string {
	// naive approach:
	// we keep track of the characters in the string with a map
	// if a duplicate occurance appears we then run through the characters
	// from the 1st occurance to the 2nd occurance
	// compare the string with reversed string
	// if equal then add as max palindrome
	// if not a palindrome we can check the sub string for any palindromes (this bit quite slow)
	// advance end pointer

	// the naive approach works however is quite inefficient
	// we know the problem requires this repeated part for substring palindrome checking (recursion)
	// where we need to check for palindromes from start:end where we move independently the start and end
	// i.e. check substring = s[start:end] where start 1...n AND check substring = s[start:end] where end 1...n
	// its clear that this problem can be broken into 2 parts so we can use dynamic programming (dp)
	// note: where we can use recursion we can optimize it with dp

	start := 0
	l := len(s)

	// store the position of the characters occurance
	chars := make(map[rune]int, l)

	palindrome := ""

	for end := 0; end < l; end++ {
		r := []rune(s)[end]
		fmt.Println("rune: ", string(r), "map: ", chars)
		pos, inMap := chars[r]

		if inMap {
			subStr := s[pos : end+1]
			fmt.Println("sub string: ", subStr)

			reversed := reverse(subStr)

			// Remove all characters before the new
			// for i := start; i <= pos; i++ {
			// 	delete(chars, r)
			// }

			if reversed == subStr && len(subStr) > len(palindrome) {
				fmt.Println("palindrome found: ", subStr)
				palindrome = subStr
			} else if reversed != subStr {
				// TODO: we want to check sub strings that match in here
				// i.e. aaca and acaa
				// start = pos + 1
				fmt.Printf("do this! s: %d, e: %d \n", start, end)
				for j := start; j < end; j++ {
					sub := s[j : end+1]
					rev := reverse(sub)
					fmt.Println("huh?: ", sub)
					if rev == sub && len(sub) > len(palindrome) {
						fmt.Println("found palindrome in sub: ", sub)
						palindrome = sub
					}
				}
			}
		} else {
			// add char => int to map
			chars[r] = end
		}
	}
	if palindrome == "" {
		palindrome = string(s[l-1])
	}

	return palindrome
}

// stackoverflow solution as im lazy
func reverse(s string) string {
	// Get Unicode code points.
	n := 0
	rune := make([]rune, len(s))
	for _, r := range s {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	// Convert back to UTF-8.
	output := string(rune)
	fmt.Println(output)
	return output
}

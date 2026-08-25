package maxvowelsinsubstring

/*
Problem: Max Number of Vowels in a Substring of Given Length

Given a string s and an integer k, find the maximum number of vowels (a, e, i, o, u) in any substring of s with length k.

Example:

s = "abciiidef", k = 3

Output: 3 (the substring "iii" has 3 vowels)

Example 2:

s = "leetcode", k = 3

Output: 2 (e.g. "lee", "eet", "ode" — a few windows tie at 2)

Constraints to think about:

1 <= s.length
1 <= k <= s.length
s consists of lowercase English letters
*/

func maxVowels(s string, k int) int {
	if k > len(s) {
		return 0
	}

	vowels := map[byte]struct{}{'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {}}
	var sum int
	// establish initial window
	for i := range k {
		if _, ok := vowels[s[i]]; ok {
			sum++
		}
	}

	maxSum := sum
	for i := range s[k:] {
		if _, ok := vowels[s[i+k]]; ok {
			sum += 1
		}

		if _, ok := vowels[s[i]]; ok {
			sum -= 1
		}

		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}

func maxVowelsRunes(s string, k int) int {
	rs := []rune(s)
	if k > len(rs) {
		return 0
	}

	vowels := map[rune]struct{}{'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {}}
	var sum int
	// establish initial window
	for i := range k {
		if _, ok := vowels[rs[i]]; ok {
			sum++
		}
	}

	maxSum := sum
	for i, r := range rs[k:] {
		var add, subtract int
		if _, ok := vowels[r]; ok {
			add = 1
		}

		if _, ok := vowels[rs[i]]; ok {
			subtract = 1
		}

		sum = sum + add - subtract

		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}

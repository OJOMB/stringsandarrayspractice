package maxsubstringwithkdistinctchars

/*
Longest Substring with At Most K Distinct Characters — classic dynamic (variable-size) sliding window since the window grows and shrinks based on a condition rather than staying a fixed length.

# Problem

Given a string s and an integer k, find the length of the longest substring that contains at most k distinct characters.

Examples:

Input: s = "eceba", k = 2
Output: 3   // "ece"

Input: s = "aa", k = 1
Output: 2   // "aa"

Input: s = "abcabcabc", k = 2
Output: 2   // e.g. "ab", "bc", etc. — never can hold 3 distinct with k=2

Constraints:

1 <= len(s) <= 5*10^4
0 <= k <= 500
s consists of lowercase/uppercase English letters (or general runes if you want to generalize)
*/

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	if len(s) < 1 || k < 1 {
		return 0
	}

	rs := []rune(s)
	// maxLen keeps track of the length of the longest valid window seen so far
	maxLen := 0
	// freqs keeps track of the frequency of each character in the current l-r window
	freqs := make(map[rune]uint)

	l := 0
	for r := range rs {
		ru := rs[r]
		_, ok := freqs[ru]
		if !ok {
			// if the character is not already in the map, initialize its frequency to 1
			freqs[ru] = 1
		} else {
			// if the character is already in the map, increment its frequency
			freqs[ru]++
		}

		// if the number of distinct characters in the current window exceeds k, shrink the window from the left until it doesn't
		for len(freqs) > k {
			lru := rs[l]
			// decrement the frequency of the leftmost character
			freqs[lru]--
			// if its frequency becomes 0, remove it from the map
			if freqs[lru] == 0 {
				delete(freqs, lru)
			}

			l++
		}

		if currLen := r - l + 1; currLen > maxLen {
			maxLen = currLen
		}
	}

	return maxLen
}

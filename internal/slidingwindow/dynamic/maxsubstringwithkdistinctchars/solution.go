package maxsubstringwithkdistinctchars

/*
Here's a solid one: Longest Substring with At Most K Distinct Characters — classic dynamic (variable-size) sliding window since the window grows and shrinks based on a condition rather than staying a fixed length.

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
	maxLen := 0
	freqs := make(map[rune]uint)
	l := 0
	for r := range rs {
		ru := rs[r]
		_, ok := freqs[ru]
		if !ok {
			freqs[ru] = 1
		} else {
			freqs[ru]++
		}

		for len(freqs) > k {
			lru := rs[l]
			freqs[lru]--
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

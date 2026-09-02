package longestrepeatingcharacterreplacement

/*

Problem

Given a string s consisting of uppercase English letters and an integer k, you can choose at most k characters in the string
and change each to any other uppercase letter. Return the length of the longest substring containing the same letter,
after performing these replacements.

Examples:

Input: s = "ABAB", k = 2
Output: 4   // replace both A's or both B's

Input: s = "AABABBA", k = 1
Output: 4   // replace one 'A' at index 3 → "AABBBBA" wait, actually:
            // window "ABBB" (indices 3-6) with the 'A' at index 3
            // replaced with 'B' works, giving 4 consecutive same letters

Constraints:

1 <= len(s) <= 10^5
s consists only of uppercase English letters
0 <= k <= len(s)
*/

func characterReplacement(s string, k int) int {
	rs := []rune(s)

	// freqs must represent what is in the current l-r window
	freqs := make(map[rune]int)
	l := 0
	// maxLen keeps track of the length of the longest valid window seen so far
	maxLen := 0
	// maxFreq keeps track of the highest freq of any character in the current l-r window
	maxFreq := 0

	for r, ru := range rs {
		freqs[ru]++
		if freqs[ru] > maxFreq {
			maxFreq = freqs[ru]
		}

		// if the lenght of the window minus the maxFreq is greater than k
		// we don't have enough swap capacity so we need to shrink the window (increment l)
		if (r-l+1)-maxFreq > k {
			lru := rs[l]
			freqs[lru]--
			l++
		}

		if currLen := r - l + 1; currLen > maxLen {
			maxLen = currLen
		}
	}

	return maxLen
}
